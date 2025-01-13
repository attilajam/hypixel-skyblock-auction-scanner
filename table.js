document.getElementById("query").value =
  new URLSearchParams(window.location.search).get("query") || "";
function copyToClipboard(uuid, button) {
  const textToCopy = `/viewauction ${uuid}`;

  navigator.clipboard
    .writeText(textToCopy)
    .then(() => {
      button.textContent = "Copied!";
      button.classList.add("copy-success");

      setTimeout(() => {
        button.textContent = "Copy";
        button.classList.remove("copy-success");
      }, 1000);
    })
    .catch((err) => {
      console.error("Failed to copy:", err);
      button.textContent = "Failed!";
      setTimeout(() => {
        button.textContent = "Copy";
      }, 1000);
    });
}
