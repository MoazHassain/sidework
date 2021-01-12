const template = document.createElement("template");
template.innerHTML = `
    <style>
        .user-card
        {
            margin: 1em;
            display: grid;
            grid-template-columns: 20% 50%;

        }
        h3
        {
            color: coral;
        }
        img
        {
            width: 80%;
            height: auto ;
            
        }
    
    </style>

    <div class="user-card">
        <img/ alt="failed to load img">
        <div>
            <h3></h3>
            <div class="info">
                <p>phone</p>
                <p>email</p>
            </div>
            <button id="toggle-info" >hide info</button>
        </div>
    </div>




`


class userCard extends HTMLElement {
    constructor() {
        super();

        this.attachShadow({ mode:"open"});
        this.shadowRoot.appendChild(template.content.cloneNode(true));
        this.shadowRoot.querySelector("h3").innerText=this.getAttribute("name");
        this.shadowRoot.querySelector("img").src=this.getAttribute("avater");
    }
}

window.customElements.define("user-card", userCard);