// Simulated _mNDetails object
window._mNDetails = window._mNDetails || {};

// Mock loadTag function
window._mNDetails.loadTag = function(slotId, size, divId) {
  // Find the target div
  const div = document.getElementById(divId);
  if (!div) {
    console.warn(`Div with ID ${divId} not found`);
    return;
  }

  // Parse size
  const [width, height] = size.split("x").map(Number);

  // Create a placeholder ad (simulating Media.net creative)
  const ad = document.createElement("div");
  ad.style.width = width + "px";
  ad.style.height = height + "px";
  ad.style.backgroundColor = "#ccc";
  ad.style.display = "flex";
  ad.style.alignItems = "center";
  ad.style.justifyContent = "center";
  ad.style.border = "1px solid #888";
  ad.style.fontFamily = "Arial, sans-serif";
  ad.style.color = "#333";
  ad.innerText = `Ad Slot: ${slotId} (${size})`;

  // Insert the ad into the div
  div.appendChild(ad);

  // Simulate tracking
  console.log(`Ad rendered in div ${divId} (Slot: ${slotId}, Size: ${size})`);
};
