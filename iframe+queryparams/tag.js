// way 1 of communicating div and server
//tag.js has publisher page context
//running a this script at publisher page to gets its publisher info
//then creates iframe and append it as child to the div tag in publisher.
//inside iframe src i am sending publisher data to ad.html


(function () {
  // const slotId = "ad-slot-1";
  // const el = document.getElementById(slotId);

  // if (!el) {
  //   console.warn("Ad slot not found:", slotId);
  //   return;
  // }

  // const adId = el.getAttribute("data-ad-id");

  // if (!adId) {
  //   console.warn("data-ad-id is missing on ad slot");
  //   return;
  // }











//generate publisher info as query params to send to iframe.
function getPublisherParams(adSlotId, adId) {
  const safe = (v, fallback = "unknown") =>
    v !== undefined && v !== null ? v : fallback;

  const data = {
    publisherid:window.publisherId,
    // --- Navigator properties ---
    ua: navigator.userAgent,
    language: navigator.language,
    platform: navigator.platform,
    cookieEnabled: navigator.cookieEnabled,
    online: navigator.onLine,
    deviceMemory: safe(navigator.deviceMemory),
    cpuCores: safe(navigator.hardwareConcurrency),
    maxTouchPoints: navigator.maxTouchPoints ?? 0,
    connectionDownlink: safe(navigator.connection?.downlink),

    // --- Window/Screen properties ---
    viewportWidth: window.innerWidth,
    viewportHeight: window.innerHeight,
    scrollX: window.scrollX,
    scrollY: window.scrollY,
    screenWidth: screen.width,
    screenHeight: screen.height,
    availWidth: screen.availWidth,
    availHeight: screen.availHeight,
    colorDepth: screen.colorDepth,
    pixelDepth: screen.pixelDepth,
    timezone: Intl.DateTimeFormat().resolvedOptions().timeZone,

    // --- Page/Ad properties ---
    adSlotId: adSlotId,
    adId: adId,
    url: window.location.href,
    ref: document.referrer || ""
  };

  return new URLSearchParams(data);
}




 




//iterating over only visible ad slots in viewport of publishereach adslots in PP for mnet
function visibledivs(){
  const cfg = window.myAdConfig || {};
  console.log(cfg.visibleAdSlots,cfg);
  (cfg.visibleAdSlots || []).forEach(adSlot=>{ 
     
  const adSlotId = adSlot.id;
 
  const el = document.getElementById(adSlotId);
  if (!el) return console.warn("Ad slot not found:", adSlotId);

   // ✅ Prevent duplicate iframes
  if (el.querySelector('iframe')) {
      console.log("Iframe already exists for", adSlotId);
      return;
    }

  const params=getPublisherParams(adSlotId,'ad-456');

  // Create iframe ,and sending publisher info using url.
  const iframe = document.createElement("iframe");
  iframe.src = "http://localhost:5501/iframe+queryparams/ad.html?" + params.toString();
  iframe.width = el.offsetWidth || 300;   // fallback width
  iframe.height = el.offsetHeight || 250; // fallback height
  iframe.style.border = "0";
  iframe.loading = "lazy";

  el.appendChild(iframe);


});
}
visibledivs();
//visibledivs function was getting Window.myadconfig.visibleadslots as undefined for intitial call
//so added a lil delay so it gets updated for initial render.
// add listerner to run visibledivs function on scroll as Window.myadconfig.visibleadslots array willbe updated.
// setTimeout(visibledivs,1);
// window.addEventListener('scroll',visibledivs);






//✅ Listen to publisher updates
  window.addEventListener('visibleAdSlotsUpdated',visibledivs);



  // // Run renderAds once DOM is fully loaded
  // if (document.readyState === 'loading') {
  //   // DOM not ready, wait for it
  //   document.addEventListener('DOMContentLoaded', visibledivs);
  // } else {
  //   // DOM already ready
  //   visibledivs();
  // }




// const params = getPublisherParams("slot-123", "ad-456");

//   console.log(params);
//   console.log(window);
//   console.log(navigator);
//   console.log(document);

  
//   // Create iframe ,and sending publisher info using url.
//   // const iframe = document.createElement("iframe");
//   // iframe.src = "http://localhost:5501/iframe+queryparams/ad.html?" + params.toString();
//   // iframe.width = el.offsetWidth || 300;   // fallback width
//   // iframe.height = el.offsetHeight || 250; // fallback height
//   // iframe.style.border = "0";
//   // iframe.loading = "lazy";

//   // el.appendChild(iframe);
})



();
