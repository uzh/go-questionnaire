/*
    change or input for a input[type=range]
    updating   the corresponding 'display'
    unchecking the corresponding 'noanswer'
    updating   the corresponding 'hidden' input;
    if the range-input is dirty;
    otherwise we just set the 'display' to '--'

    TODO
    these funcs make use of relative HTML node positions
    they should be simplified by exploiting the ID correspondance

        let catcherID  = src.id.replace("_noanswer", "_hidd");
        let displayID  = src.id +  "_display");
        let noanswerID = src.id +  "_noanswer");


*/
function pdsRangeInput(src) {
    // console.log("rangeInput()");

    // src.style.backgroundColor = "transparent";

    if (!src.parentNode) {
        return true
    }
    if (!src.parentNode.childNodes) {
        return true
    }

    let chn = src.parentNode.childNodes;
    // console.log("child nodes num", chn.length);

    let noAnsw = null;
    let label = null;
    for (i = 0; i < chn.length; i++) {

        let el = chn[i];

        if (el.nodeType == Node.TEXT_NODE) {
            // console.log("   ch #",i , " - textnode");
        } else {
            // console.log("   ch #", i, el.nodeType, el.type, el.name);
            if (el.type == "radio") {
                noAnsw = el;
            }
            if (el.tagName == "LABEL") {
                label = el;
            }
        }
    }

    if (noAnsw) {
        noAnsw.checked = false;
        // rg.disabled = true;
    }
    if (label) {

        let chn = label.childNodes;
        // console.log("  label child nodes num", chn.length);

        let display = null;
        for (i = 0; i < chn.length; i++) {

            let el = chn[i];

            if (el.nodeType == Node.TEXT_NODE) {
                // console.log("   ch #",i , " - textnode");
            } else {
                // nodeType is an integer
                // console.log("       ch #", i, el.nodeType, el.tagName, el.type, el.name);
                if (el.type == "text") {
                    display = el;
                }
            }
        }

        if (display) {

            // console.log("src.value=", src.value, "data-dirty:", src.dataset.dirty);

            let incr = parseFloat(src.value) + parseFloat(src.step);
            let out = ""
            if (src.value) {
                out += src.value;
            }
            out += " - ";
            out += incr;
            display.value = out;

            let catcher = document.getElementById(src.id + "_hidd");
            if (catcher) {
                catcher.value = src.value;
            }

        }

    }
    return true;
}


// activate an input[type=range] from de-activated visual state 
function pdsRangeClick(src) {
    // console.log("rangeClick()");
    // src.style.backgroundColor = "transparent";
    src.classList.remove("hidethumb");
    src.classList.remove("noanswer");
    src.dataset.dirty = "";

    // src.disabled = false;
    return true;
}

// 'no answer' radio puts corresponding range-input into de-activated state
function pdsRangeRadioInput(src) {
    // console.log("rangeRadioInput()");
    // console.log("rangeRadioInput() - src id", src.id);

    if (!src.parentNode) {
        return true
    }
    if (!src.parentNode.childNodes) {
        return true
    }

    const chn = src.parentNode.childNodes;
    // console.log("child nodes num", chn.length);

    let rg = null;
    let label = null;
    for (i = 0; i < chn.length; i++) {

        let el = chn[i];

        if (el.nodeType == Node.TEXT_NODE) {
            // console.log("   ch #",i , " - textnode");
        } else {
            // console.log("   ch #", i, el.nodeType, el.type, el.name);
            if (el.type == "range") {
                rg = el;
            }
            if (el.tagName == "LABEL") {
                label = el;
            }
        }
    }

    if (rg) {
        // rg.style.backgroundColor = "darkgray";
        rg.classList.add("hidethumb");
        rg.classList.add("noanswer");

        // rg.disabled = true;
        let catcherID = src.id
        catcherID = catcherID.replace("_noanswer", "_hidd");
        let catcher = document.getElementById(catcherID);
        if (catcher) {
            catcher.value = "no answ.";
        } else {
            console.log("catcher not found", catcherID);
        }

    }

    if (label) {

        let chn = label.childNodes;
        // console.log("  label child nodes num", chn.length);

        let display = null;
        for (i = 0; i < chn.length; i++) {

            let el = chn[i];

            if (el.nodeType == Node.TEXT_NODE) {
                // console.log("   ch #",i , " - textnode");
            } else {
                // console.log("       ch #", i, el.nodeType, el.type, el.name);
                // input[type=text]
                if (el.type == "text") {
                    display = el;
                }
            }
        }

        if (display) {
            // also adapt in questionaire-grid.go dispVal and checked
            display.value = "no answ.";
        }

    }


    return true;
}