const slot = "ad-slot-1";
const adId = 123;

const iframe = document.createElement("iframe");
iframe.src = "http://localhost:5500/iframe+postmessage/ad.html";
iframe.width = 300;
iframe.height = 250;
iframe.style.border = "0";

document.getElementById(slot).appendChild(iframe);

iframe.onload = () => {
  iframe.contentWindow.postMessage({
    slot,
    adId,
    pageUrl: window.location.href,
    referrer: document.referrer,
    userAgent: navigator.userAgent
  }, "*");
};
