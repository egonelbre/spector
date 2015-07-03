function RenderTimeline(context, timeline, size){
	var processHeaderHeight = 20;
	var threadHeaderHeight = 10;
	var layerHeight = 20;

	context.save();

	var left = 0, right = size.x;

	var start = timeline.view.start, end = timeline.view.end;
	var scale = (right - left) / (end - start);

	var top = 0;
	timeline.view.top = 0;

	var stripe = 0;
	timeline.processes.map(function(process){
		process.view.top = top;

		context.fillStyle = "#fff";
		context.fillRect(left, top, right-left, processHeaderHeight);

		context.fillStyle = "#f00";
		context.fillRect(left, top, right-left, 1);
		top += 1;

		context.font = "14px monospace";
		context.fillStyle = "#000";
		context.fillText("pid=" + process.pid, left + 2, top + 14);

		top += processHeaderHeight - 2;

		context.fillStyle = "#ccc";
		context.fillRect(left, top, right-left, 1);
		top += 1;

		process.threads.map(function(thread){
			thread.view.top = top;

			context.fillStyle = "#ccc";
			context.fillRect(left, top, right-left, threadHeaderHeight);

			context.font = "8px monospace";
			context.fillStyle = "#000";
			context.fillText("tid=" + thread.tid, left + 2, top + 8);

			top += threadHeaderHeight;
			top++;

			var durations = thread.durations;
			durations.layers.map(function(layer, layerIndex){
				context.fillStyle = "#fff";
				context.fillRect(left, top, right-left, layerHeight);

				context.beginPath();

				var n = layer.count;
				for(var i = 0; i < n; i++){
					if(layer.end(i) < start){ continue; }
					if(layer.begin(i) > end){ break; }

					var ls = left + (layer.begin(i) - start) * scale;
					var rs = left + (layer.end(i) - start) * scale;

					var w = (rs-ls)|0;
					w = Math.max(w, 1);

					context.rect(ls|0, top, w, layerHeight);
				}

				context.fillStyle = "hsla(" + 30 * layerIndex + ", 70%, 60%, 1)";
				context.fill();

				top += layerHeight;
			});

			context.beginPath();
			var instants = thread.instants;
			var n = instants.count;
			for(var i = 0; i < n; i++){
				var t  = left + (instants.time(i) - start) * scale;
				context.moveTo(t, thread.view.top);
				context.lineTo(t, top);
			}

			context.strokeStyle = "#0f0";
			context.stroke();
		});

		context.beginPath();
		var instants = process.instants;
		var n = instants.count;
		for(var i = 0; i < n; i++){
			var t  = left + (instants.time(i) - start) * scale;
			context.moveTo(t, process.view.top);
			context.lineTo(t, top);
		}
		context.strokeStyle = "#f00";
		context.stroke();
	});


	context.beginPath();
	var instants = timeline.instants;
	var n = instants.count;
	for(var i = 0; i < n; i++){
		var t  = left + (instants.time(i) - start) * scale;
		context.moveTo(t, timeline.top.view);
		context.lineTo(t, top);
	}
	context.strokeStyle = "#00f";
	context.stroke();

	context.beginPath();
	var flows = timeline.flows;
	var n = flows.count;
	for(var i = 0; i < n; i++){
		if(flows.start(i) == 0){ continue; }
		if(flows.finish(i) == 0){ continue; }

		var t  = left + (instants.time(i) - start) * scale;

		var sev = timeline.events[flows.startEID(i)];
		var st = timeline.trackByID(sev.pid, sev.tid);
		var sp  = left + (flows.start(i) - start) * scale;

		var fev = timeline.events[flows.finishEID(i)];
		var ft = timeline.trackByID(fev.pid, fev.tid);
		var fp  = left + (flows.finish(i) - start) * scale;

		context.moveTo(sp, st.view.top + layerHeight);
		context.bezierCurveTo(
			(sp + fp) / 2, st.view.top + layerHeight,
			(sp + fp) / 2, ft.view.top + layerHeight,
			fp, ft.view.top + layerHeight
		);
	}

	context.lineWidth = 0.5;
	context.strokeStyle = "#888";
	context.stroke();

	context.restore();
}