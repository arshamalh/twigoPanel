<script>
    import Chart from 'chart.js/auto';
	import { onMount } from 'svelte';

	export let jsonData = {}
	let chartCanvas
	let chart

	const data = {
		labels: [],
		datasets: [
			{
				label: 'Likes',
				data: [],
				fill: false,
				borderColor: 'rgb(75, 192, 192)',
				tension: 0.1
			},
			{
				label: 'Retweets',
				data: [],
				fill: false,
				borderColor: 'rgb(255, 99, 132)',
				tension: 0.1
			},
			{
				label: 'Quotes retweets',
				data: [],
				fill: false,
				borderColor: 'rgb(255, 159, 64)',
				tension: 0.1
			},
			{
				label: 'Replies',
				data: [],
				fill: false,
				borderColor: 'rgb(255, 205, 86)',
				tension: 0.1
			}
		]
	};

	onMount(() => {
		const ctx = chartCanvas.getContext("2d");
		chart = new Chart(ctx, {
			type: 'line',
			data: data,
			options: {
				maintainAspectRatio: false,
				responsive: true,
			}
		});
		let labels = []
		let like_data = []
		let retweet_data = []
		let reply_data = []
		let quote_data = []
		let timeStamps = Object.keys(jsonData)
		let first_ts = timeStamps[0]
        for (const ts of timeStamps) {
            labels.push(Math.round((parseInt(ts) - parseInt(first_ts))/60))
            like_data.push(jsonData[ts]["like_count"])
            retweet_data.push(jsonData[ts]["retweet_count"])
            reply_data.push(jsonData[ts]["reply_count"])
            quote_data.push(jsonData[ts]["quote_count"])
        }
        chart.data.labels = labels
        chart.data.datasets[0].data = like_data
        chart.data.datasets[1].data = retweet_data
        chart.data.datasets[2].data = quote_data
        chart.data.datasets[3].data = reply_data
        chart.update()
	});
</script>

<canvas bind:this={chartCanvas}></canvas>

<style>
    canvas {
        margin-bottom: 40px;
    }
</style>
