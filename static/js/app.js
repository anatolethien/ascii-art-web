function toggle_dark() {
    const body = document.body;
    if (document.body.classList.contains("dark")) {
        document.body.classList.remove("dark");
        document.body.classList.add("light");
        document.querySelector(".darkmode-btn").innerHTML = `<img src="/static/images/dark.svg" title="Toggle dark mode">`;
        document.querySelector(".copy-btn").innerHTML = `<img src="/static/images/clipboard.svg" title="Copy to clipboard">`;
        document.querySelector(".download-btn").innerHTML = `<img src="/static/images/download.svg" title="Download as file">`;
    } else if (document.body.classList.contains("light")) {
        document.body.classList.remove("light");
        document.body.classList.add("dark");
        document.querySelector(".darkmode-btn").innerHTML = `<img src="/static/images/light.svg" title="Toggle light mode">`;
        document.querySelector(".copy-btn").innerHTML = `<img src="/static/images/clipboard-dark.svg" title="Copy to clipboard">`;
        document.querySelector(".download-btn").innerHTML = `<img src="/static/images/download-dark.svg" title="Download as file">`;
    }
}

function copy_clipboard() {
    let output = document.querySelector("pre").innerText;
    let temp = document.createElement("textarea");
    document.body.appendChild(temp);
    temp.value = output;
    temp.focus();
    temp.select();
    document.execCommand("copy");
    temp.remove();
}

function download_file() {
    let blob = new Blob([document.querySelector("pre").innerText], {
        type: "text/plain"
    });
    let url = window.URL.createObjectURL(blob);
    let temp = document.createElement("a");
    temp.href = url;
    temp.download = "ascii-art.txt";
    temp.click();
    temp.remove();
}

document.querySelector(".darkmode-btn").addEventListener("click", toggle_dark);
document.querySelector(".copy-btn").addEventListener("click", copy_clipboard);
document.querySelector(".download-btn").addEventListener("click", download_file);
