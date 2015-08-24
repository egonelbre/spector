package("spector.import.simulator", function(){
	depends("/spector/stream.js");
	depends("/spector/protocol.js");

	function R(probability){ return Math.random() < probability; }

	function len(obj){
		if(obj instanceof Array) { return obj.length; }
		var c = 0;
		for(var name in obj) { if(obj.hasOwnProperty(name)){ c++; } }
		return c;
	}

	function pick(obj){
		if(obj instanceof Array) { return (Math.random()*obj.length)|0; }
		var i = Math.random() * len(obj);
		for(var name in obj){
			if(obj.hasOwnProperty(name)){ if(i <= 0){ return name; } i--; }
		}
		for(var name in obj){ return name; }
		return undefined;
	}

	var Event = spector.Event;
	function Stream(props){
		props = props | {};
		var maxOpenEvents  = props.maxOpenEvents  || 8;
		var maxAsyncEvents = props.maxAsyncEvents || 2;
		var maxThreads     = props.maxThreads     || 20;
		var emitInterval   = props.emitInterval   || 100;
		var threadInterval = props.threadInterval || 150;

		this.time_ = 0;
		this.threads_ = {};
		this.buffer_ = [];
		this.stage = spector.Stream.Stage.Reading;

		this.lastThreadID = 0;
		this.lastEventID = 0;

		var stream = this;

		function Thread(){
			this.TID = stream.lastThreadID++;
			this.open_ = {};
			this.interval_ = undefined;
		}
		Thread.prototype = {
			emit_: function(){
				stream.advanceTime_();

				var count = len(this.open_);
				if(R(count/maxOpenEvents)){
					var eid = pick(this.open_);
					delete(this.open_[eid]);
					stream.emit(new Event.End({
						Time: stream.time_,
						ThreadID: this.TID,
						StackID: 0,
						ID: parseInt(eid)
					}));
				} else {
					var eid = stream.lastEventID++;
					this.open_[eid] = eid;
					stream.emit(new Event.Begin({
						Time: stream.time_,
						ThreadID: this.TID,
						StackID: 0,
						ID: eid
					}));
				}
			},
			start_: function(){
				this.interval_ = window.setInterval(
					this.emit_.bind(this),
					threadInterval
				);
				stream.emit(new Event.ThreadStart({
					Time: stream.time_,
					ThreadID: this.TID,
					StackID: 0,
				}));
			},
			stop_: function(){
				window.clearInterval(this.interval_);
				this.interval_ = undefined;
				stream.emit(new Event.ThreadStop({
					Time: stream.time_,
					ThreadID: this.TID,
					StackID: 0,
				}));
			},
			kill: function(){ window.clearInterval(this.interval_); }
		};

		this.advanceTime_ = function(){
			this.time_ += (Math.random() * 100) | 0;
		};

		this.emit_ = function(){
			this.advanceTime_();
			var count = len(this.threads_);
			if(R(count/maxThreads)){
				var tid = pick(this.threads_);
				var thread = this.threads_[tid];
				delete(this.threads_[tid]);
				thread.stop_();
			} else {
				var thread = new Thread();
				this.threads_[thread.TID] = thread;
				thread.start_();
			}
		};

		this.stop_ = function(){
			for(var tid in this.threads_) {
				this.threads_[tid].kill();
			}
			window.clearInterval(this.interval_);
			this.interval_ = undefined;
			this.emit(new Event.StreamStop({ Time: this.time }));
		};

		this.emit = function(ev){
			this.buffer_.push(ev);
		};

		this.close = function(){
			this.stop_();
		};

		this.next = function(){
			if(this.buffer_.length <= 0) {
				return undefined;
			}
			return this.buffer_.shift();
		};

		this.emit(new Event.StreamStart({
			ProcessID: Math.random() * (1 << 30) | 0,
			MachineID: Math.random() * (1 << 30) | 0,
			Time: this.time,
			CPUFrequency: 1e9 // ticks per second
		}));

		this.meta = {};

		this.interval_ = window.setInterval(this.emit_.bind(this), threadInterval);
	};

	return {
		Stream: Stream
	};
});
</script>