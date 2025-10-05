const form = document.getElementById("form-data");
const outputContainer = document.getElementById("result-container");
// const output = document.getElementById("result");
const output = document.createElement("a");

form.addEventListener("submit", async (e) => {
  e.preventDefault();
  const formData = new FormData(form);
  const url = formData.get("url");

  try {
    const res = await fetch(form.action, {
      method: form.method,
      body: new URLSearchParams({ url }),
    });
    const hash = await res.text();
    const link = `${window.location.origin}/${hash}`;
    output.href = link;
    output.textContent = link;
    output.target = "_blank";
    output.setAttribute("style", "text-decoration: none; color: #01BFCF;");
    outputContainer.textContent = `Short url: `;
    outputContainer.appendChild(output);
  } catch (err) {
    console.log(err);
  }
});
