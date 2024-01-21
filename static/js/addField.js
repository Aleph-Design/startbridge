
let commisCnt = 0;

function addField(plusEl) {
  // get a reference to grandfather node (member-stack | commis-stack)
  const parentEl = plusEl.parentElement.parentElement;

  if (plusEl.previousElementSibling.value.trim() === "") {
    // lblEl = parentEl.getElementsByTagName("label")[0];
    // lblEl.style.color = "#dd0000";
    // lblEl.innerHTML = "Je bent leeg vriend.";

    // plusEl.previousElementSibling.style.border = "2px solid #880000";
    // plusEl.previousElementSibling.style.color = "#880000";
    // plusEl.previousElementSibling.value = "Eerst een naam invullen."
    return false;
  }
  // lblEl = parentEl.getElementsByTagName("label")[0];
  // lblEl.style.color = "#000000";
  // lblEl.innerHTML = "";
  // create a separated innerHTML for each form-group; also a separate
  // form.value for each appearance  such as: "board-2 = [Banaan]"
  switch (parentEl.id) {
    case "member-stack":
      newNodes = `
      <input type="text" class="form-control" name="member" placeholder="Type drie letters.">`
      console.log("newNodes: ", newNodes)
    break;
    case "commis-stack":
      ++commisCnt;
      newNodes = `
      <input type="text" class="form-control" name="commis-${commisCnt}" placeholder="Type drie letters">`
    break;
  }
  newNodes += `
    <span onClick="addField(this)">+</span>
    <span onClick="minField(this)">-</span>
  `
  // create a div container for the class="field"
  const divNode = document.createElement('div');
  divNode.setAttribute("class", "field");
  divNode.innerHTML = newNodes
  parentEl.append(divNode)
  // Unhide the minus sign
  plusEl.nextElementSibling.style.display = "inline-block";
  // Hide the plus sign
  plusEl.style.display = "none";
}

function minField(minEl) {
  minEl.parentElement.remove();
}

// on submit we want all separate form.values arranged in one array
const form = document.forms[0];
form.addEventListener('submit', (e) => {
  // prevent the form to communicate with the server.
  // this will prevent the form to submit????????????
  e.preventDefault();
  // get all? values from the input value attributes
  const formData = new FormData(form);
  // store selected values in an array
  let members = [];
  for (const [key, value] of formData) {
    if (key === "member") {
        if (value !== "") {members.push(value);}
        // now field "name=" holds the value of members array
    }
  }
  if (members !== "") {
    document.querySelector("#board-member").value = members;
  }
  console.log("members: ", members)
  // NOW, we definitly submit the form.
  form.submit();
});

/************************
 *   BOARD - SECTION    *
 ************************/
// const chrMan = document.querySelector('#chair-man');
// chrMan.addEventListener('click', (e) => {
  // search for a chairman.
// });