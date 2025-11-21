//extracting info given iframe in src.
//sending load to server.
//ad.js has context of iframe , iframe context not publisher page context.

(function () {
  
  const p = new URLSearchParams(window.location.search);

 const payload = (function getpayload(p){
  // --- Navigator properties ---
  return {publisherid:p.get("publisherid"),
    userAgent: p.get("ua"),
  language: p.get("language"),
  platform: p.get("platform"),
  cookieEnabled: p.get("cookieEnabled"),
  online: p.get("online"),
  deviceMemory: p.get("deviceMemory"),
  cpuCores: p.get("cpuCores"),
  maxTouchPoints: p.get("maxTouchPoints"),
  connectionDownlink: p.get("connectionDownlink"),

  // --- Window/Screen properties ---
  viewportWidth: p.get("viewportWidth"),
  viewportHeight: p.get("viewportHeight"),
  scrollX: p.get("scrollX"),
  scrollY: p.get("scrollY"),
  screenWidth: p.get("screenWidth"),
  screenHeight: p.get("screenHeight"),
  availWidth: p.get("availWidth"),
  availHeight: p.get("availHeight"),
  colorDepth: p.get("colorDepth"),
  pixelDepth: p.get("pixelDepth"),
  timezone: p.get("timezone"),

  // --- Page/Ad properties ---
  adSlotId: p.get("adSlotId"),
  adId: p.get("adId"),
  pageUrl: p.get("url"),
  referrer: p.get("ref")};
})(p);





// in iframe window context appending slotid of div as myiframe is attached to that div.
window.adSlotId=payload.adSlotId;
  console.log("getting info from publisher inside iframe using query params.",payload);
console.log("iframe window",payload.adSlotId,window);













  // Render HTML creative
  document.body.innerHTML = `
    <div style="border:1px solid #ccc; padding:10px; text-align:center;">
      <h3>Buy Now!</h3>
      <p>Awesome product only for you.</p>
      <button id="cta-btn">Click</button>
    </div>
  `;





  // --- Simulate click ---
  const btn = document.getElementById("cta-btn");
  if (btn) {
    btn.addEventListener("click", () => {
      alert(`🟡 Ad Clicked!\nSlot ID: ${window.adSlotId}\nAd ID: ${payload.adId}`);
    });
  }



  
})();
