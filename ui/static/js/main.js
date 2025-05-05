document.addEventListener("DOMContentLoaded", function() {
    document.body.addEventListener("htmx:beforeSwap", function(evt) {
        evt.preventDefault();
        const hash = evt.detail.xhr.response;
        document.querySelector("#result").innerHTML =
            `<p>Short URL: <a href="http://localhost:8080/${hash}">http://localhost:8080/${hash}</a></p>`;
    });
});
