window.addEventListener("message", e => {
  const pubInfo = e.data;
  console.log("Publisher info received using iframe+postmessage.:", pubInfo);

  // Send impression
// fetch("https://ads.example.com/impression", {
//     method: "POST",
//     headers: { "Content-Type": "application/json" },
//     body: JSON.stringify(pubInfo)
//   });

  // Render ad
  document.body.innerHTML = `<h3>Buy Now!</h3><button onclick="alert('Clicked')">Click Me</button>`;
});
