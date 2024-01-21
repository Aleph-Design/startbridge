
const home = document.querySelector('#home');
home.addEventListener('click', (e) => {
  const id = e.target.previousElementSibling.innerText; // 1, 2, 3, 5, 7, etc
  const cl = e.target.className;  // pers | coop
  console.log("id: ", id);
  console.log("class: ", cl);
  /*
  zoek nu in de database to return the correct respons
  this is only just a persoon OR an organisation
  if (class == pers) {
    een persoon met id = id
  } else if (class == coop) {
    een cooperatie met id = id
  }
  */
})