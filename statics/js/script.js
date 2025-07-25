const cards = document.querySelectorAll(".artistCard");
const searchInput = document.querySelector("#search");
const parent = document.querySelector(".parent");
const allCards = Array.from(document.querySelectorAll(".artistCard"));

OnPressKeyBoard();

searchInput.addEventListener("keyup", OnPressKeyBoard);

function OnPressKeyBoard() {
    const query = searchInput.value.toLowerCase().trim();
    parent.innerHTML = ""; 
    allCards.forEach(card => {
        const name = card.id.toLowerCase();
        if (query === "" || name.includes(query)) {
            parent.appendChild(card);
        }
    });
}