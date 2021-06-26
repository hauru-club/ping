const packetsDiv = document.getElementById("packets");

const evtSource = new EventSource("/events/test");

evtSource.addEventListener("message", (event) => {
  let parsed = JSON.parse(event.data);
  let now = new Date();
  let time = now.toLocaleString("en-US", {
    year: "2-digit",
    month: "2-digit",
    day: "2-digit",
  });
  time += " ";
  time += now.toLocaleTimeString("en-US", {
    hour12: false,
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
  });
  packetsDiv.innerHTML += `<p>
    ${time}
    <b>PING</b>: <b>src</b>=${parsed.src}
    <b>dst</b>=${parsed.dst} <b>seq</b>=${parsed.seq}
    <b>len</b>=${parsed.len}
  </p>`;
  packetsDiv.scrollTo(0, packetsDiv.scrollHeight);
});
