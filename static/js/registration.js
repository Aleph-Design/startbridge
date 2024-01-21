
const checkBox = document.querySelector('#check-box');
const lblText = document.querySelector('.form-check-label');
const brdColumn = document.querySelector('#board-column');
const btnSubmit = document.querySelector('btn-submit');
checkBox.addEventListener('click', (e) => {
  if (checkBox.checked) {
    checkBox.nextElementSibling.style.background = "#ff8c00";
    lblText.innerHTML = "Bestuursleden ACTIEF";
    brdColumn.classList.remove("dim-board-column");
  } else {
    checkBox.nextElementSibling.style.background = "rgb(36, 83, 83)";
    lblText.innerHTML = "Bestuursleden IN-actief";
    brdColumn.classList.add("dim-board-column");
  }
});

window.addEventListener('load', (e) => {
  if (checkBox.checked) {
    brdColumn.classList.remove("dim-board-column");
  } else {
    brdColumn.classList.add("dim-board-column");
  }
});

function getInput(elem, csrfToken) {
  if (elem.value.length === 3) {
    const chars = /^[A-Za-z]+$/;
    const elemId = elem.id
    let val = elem.value
    if (val.match(chars)) {
      elem.value = val[0].toUpperCase() + val.slice(1).toLowerCase();
      searchRecord(elem, csrfToken)  
    } else {
      elem.value = "Alleen letters!";
      val='';
    }
  }
}

const searchRecord = async (elem, csrfToken) => {
  const formData = new FormData();
  formData.set("csrf_token", csrfToken); 
  formData.set("dataTable", elem.name);  // table to search through
  formData.set("searchArg", elem.value); // first three characters

  try {
    const response = await fetch('/post-search-json', {
      method: "post",
      body: formData,
    });
    if (!response.ok) {
      throw new Error('Request failed');
    }
    const data = await response.json();
    console.log("data: ", data)
    selectData(elem, data.items)
  } catch (error) {
    console.log(error);
  }
};


function selectData(elem, dataItems) {
  console.log("elem: ", elem)
  // Display data
  let lst = '';
  switch (elem.id) {
    case 'province-id':
      lst = document.querySelector('#province-list');
      break;
    // case 'county':
    //   lst = document.querySelector('#list-county');
    //   break;
    default:
      console.log(`No more ${elem.id}`)
  }
  let list = "";
  if (dataItems && dataItems !== "null" && dataItems !== "undefined") {
    for(i=0; i<dataItems.length; i++) {
      list += '<li>' + dataItems[i].id + ' - ' + dataItems[i].title + '</li>';
    }
  } else {
    list += '<li>' + "Bestaat NIET!" + '</li>';
  }
  console.log("list:", list)
  lst.innerHTML = '<ul>' + list + '</ul>';
  lst.style.display = "block";

  // Select data
  function onClick(e) {
    // populate the input elements value attribute
    elem.value = e.target.innerText   // "Drente" oid
    // clear the list & remove dropdown
    lst.innerHTML = '';
    lst.style.display = 'none';
  }
  lst.addEventListener('click', onClick);
}