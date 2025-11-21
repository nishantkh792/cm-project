(function() {
  console.log("Shadow DOM Ad tag loaded");

  const adSlots = document.querySelectorAll('.my-ad-slot');

  adSlots.forEach((slot, index) => {
    // Create shadow root
    const shadow = slot.attachShadow({ mode: 'open' });

    // Create ad container
    const adWrapper = document.createElement('div');
    adWrapper.innerHTML = `
      <style>
      body{
      background:blue;
      color:pink;
      
      width: 100%;
      height: 100%;}

        .ad-box {
          font-family: Arial, sans-serif;
          background: black;
          color:white;
          padding: 10px;
          border: 1px solid #ddd;
          text-align: center;
          animation: fadeIn 0.5s ease;
        }
        .ad-box button {
          background: black;
          color: white;
          border: none;
          padding: 6px 10px;
          cursor: pointer;
          border-radius: 4px;
        }
        @keyframes fadeIn {
          from { opacity: 0; }
          to { opacity: 1; }
        }
      </style>
      <body>

      <div class="ad-box">
        <h4>Ad #${index + 1}</h4>
        <p>Buy your favorite gadget today!</p>
        <button id="buyBtn">Shop Now</button>
    
      </div>
      
      </body>
    `;

    shadow.appendChild(adWrapper);

    // JS inside shadow still has full context access
    const btn = shadow.querySelector('#buyBtn');
    btn.addEventListener('click', () => {
      alert("This ad JS runs in the SAME window context");
    });

    // 🔍 Demonstrate HTML encapsulation
    console.log("Publisher document.querySelector('#buyBtn') →", document.querySelector('#buyBtn'));
    console.log("Shadow root querySelector('#buyBtn') →", shadow.querySelector('#buyBtn'));
  });
})();
