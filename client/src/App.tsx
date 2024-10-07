import './App.css';

function App() {
	async function sendRequest() {
		let res = await fetch('/api');

		let data = await res.json();

		console.log(data);
	}
	return (
		<>
			<button onClick={() => sendRequest()}>send</button>
		</>
	);
}

export default App;
