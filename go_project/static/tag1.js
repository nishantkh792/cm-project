//contract is here

window.showAdtag = function (adSlotId, size) {

    console.log('publisherurl',window.location.ancestorOrigins);
    console.log(document.title,'title');
    function getPublisherParams(adSlotId, adId) {
       

        const data = {
            // --- Page/Ad properties ---
            adSlotId: adSlotId,
            adId: adId,
            publisherUrl: window.location.ancestorOrigins  [window.location.ancestorOrigins.length-1] || window.location.href,
            referrer: document.referrer || "",
            tsize:size,
            publisherPageTitle:document.title||"",
            publisherid: window.publisherId || 1234,
            

            // --- Navigator properties ---
            ua: navigator.userAgent,
            language: navigator.language,
            

            // --- Window/Screen properties ---
            viewportWidth: window.innerWidth,
            viewportHeight: window.innerHeight,
            scrollX: window.scrollX,
            scrollY: window.scrollY,
            screenWidth: screen.width,
            screenHeight: screen.height,
           
            timezone: Intl.DateTimeFormat().resolvedOptions().timeZone,

            
        };

        return new URLSearchParams(data);
    }
    const el = document.getElementById(adSlotId);
    if (!el) return console.warn("Ad slot not found:", adSlotId);

    const params = getPublisherParams(adSlotId, 'ad-456');

    // Create iframe ,and sending publisher info using url.
    const iframe = document.createElement("iframe");
    iframe.src = "http://localhost:8080/server/ad.html?" + params.toString();
    iframe.width = Number(size.split("x")[0]) || 200;   // fallback width
    iframe.height = Number(size.split("x")[1])  || 200; // fallback height
    iframe.style.border = "0";


    el.appendChild(iframe);

};



const cfg = window.myAdConfig || {};
console.log("AdSlots Queue:", cfg.AdSlots);

// Loop over queued items and execute them
(cfg.AdSlots || []).forEach(fn => {

    if (typeof fn !== "function") {
        console.warn("AdSlot entry is not a function:", fn);
        return;
    }

    console.log("Running queued slot function…");
    fn();   // <---- ONLY CALLS showAtag() INSIDE function
});
