<script>
	import Chart from './Chart.svelte';

	let files = []

	// function setData(event) {
	// 	const fileReader = new FileReader();
	// 	for (const file of event.target?.files) {
	// 		fileReader.readAsText(file)
	// 		fileReader.onload = (e) => {
	// 			jsonData = JSON.parse(e.target.result).public_metrics
	// 			chart.data = jsonData
	// 		}
	// 	}
	// }

	function extractJSONDataFromFile(file) {
		return new Promise((resolve, reject) => {
			const fileReader = new FileReader();
			fileReader.readAsText(file)
			fileReader.onload = (e) => {
				resolve(JSON.parse(e.target.result))
			}
		})
	}

	function author_link(file_name, json_data) {
		let author_id;
		if (!!json_data["author_id"]) {
			author_id = json_data["author_id"]
		} else {
			author_id = file_name.split("_")[4].split(".")[0]
		}

		return `https://twitter.com/i/user/${author_id}`
	}

	function tweet_link(file_name, json_data) {
		let tweet_id;
		if (!!json_data["id"]) {
			tweet_id = json_data["id"]
		} else {
			tweet_id = file_name.split("_")[1]
		}

		return `https://twitter.com/i/web/status/${tweet_id}`
	}

	function make_time(timestamp) {
		return 
	}
</script>

<main>
	<input type="file" multiple bind:files>
	{#each files as file}
		<div>
			{#await extractJSONDataFromFile(file)}
				<p>waiting...</p>
			{:then jsonData}
				<p>
					<a href="{author_link(file.name, jsonData)}" target="_blank">Author</a> - 
					<a href="{tweet_link(file.name, jsonData)}" target="_blank">Tweet</a>
					<br/>
					<small>
						{new Date(parseInt(jsonData["created_id"])*1000)}
					</small>
				</p>
				<Chart jsonData={jsonData.public_metrics}/>
			{/await}
		</div>
	{/each}
</main>

<style>
	main {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}

	div {
		width: 100%;
		height: 100vh;
		margin-top: 20px;
	}

	p {
		text-align: center;
	}

	input {
		margin-bottom: 10px;
	}
</style>