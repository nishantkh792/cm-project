
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




 



//MNET APPROACH
//iterating over only visible ad slots in viewport of publishereach adslots in PP for mnet
function visibledivs(){
  const cfg = window.myAdConfig || {};
  console.log(cfg.AdSlots,cfg);
  (cfg.AdSlots || []).forEach(adSlot=>{ 
     console.log('yo');
  const adSlotId = adSlot.id;
 console.log(adSlotId);
 
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
  iframe.src = "http://localhost:5501/test/ad.html?" + params.toString();
  iframe.width = el.offsetWidth || 300;   // fallback width
  iframe.height = el.offsetHeight || 250; // fallback height
  iframe.style.border = "0";
  

  el.appendChild(iframe);

});
}
visibledivs();











//visibledivs function was getting Window.myadconfig.visibleadslots as undefined for intitial call
//so added a lil delay so it gets updated for initial render.
// add listerner to run visibledivs function on scroll as Window.myadconfig.visibleadslots array willbe updated.
// setTimeout(visibledivs,1);
// window.addEventListener('scroll',visibledivs);






// ✅ Listen to publisher updates
  // window.addEventListener('visibleAdSlotsUpdated', (e) => {
  //   const slots = e.detail.visibleAdSlots || [];
  //   slots.forEach(slot => renderAdSlot(slot.id));
  // });



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




})();













//tag.js queryparams old folder working
// function yo() {
//   // Generate query params for iframe
//   function getPublisherParams(adSlotId, adId) {
//     const data = {
//       publisherid: window.publisherId,
//       adSlotId: adSlotId,
//       adId: adId,
//       url: window.location.href,
//       ref: document.referrer || ""
//     };
//     return new URLSearchParams(data);
//   }

//   // Keep track of already rendered divs globally
//   window.renderedDivs = window.renderedDivs || new Set();

//   // Render ads in top-to-bottom DOM order
//   window.renderAds = function () {
//     const divs = document.querySelectorAll('div[id^="mnet-tag-"]');

//     for (let i = 0; i < divs.length; i++) {
//       const div = divs[i];
//       const adSlotId = div.id;

//       // Skip already rendered divs
//       if (window.renderedDivs.has(adSlotId)) continue;

//       // Mark div as rendered
//       window.renderedDivs.add(adSlotId);

//       // Create iframe
//       const params = getPublisherParams(adSlotId, 'ad-456');
//       const iframe = document.createElement('iframe');
//       iframe.src = "http://localhost:5501/iframe+queryparams/ad.html?" + params.toString();
//       iframe.width = div.offsetWidth || 300;
//       iframe.height = div.offsetHeight || 250;
//       iframe.style.border = '0';
//       iframe.loading = 'lazy';

//       div.appendChild(iframe);
//     }
//   };

//   // Run once DOM is ready
//   if (document.readyState === 'interactive') {
//     document.addEventListener('DOMContentLoaded', window.renderAds);
//   } else {
//     window.renderAds();
//   }

//   // Poll the DOM every 500ms to detect newly added divs
//   window.setInterval(window.renderAds, 500);

// }
