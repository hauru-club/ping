const packetsDiv = document.getElementById("packets");

const evtSource = new EventSource("/events/test");

evtSource.addEventListener('message', (event) => {
  let parsed = JSON.parse(event.data);
  packetsDiv.innerHTML += `<p>
    <b>PING</b>: <b>src</b>=${parsed.src}
    <b>dst</b>=${parsed.dst} <b>seq</b>=${parsed.seq}
    <b>len</b>=${parsed.len}
  </p>`
});
