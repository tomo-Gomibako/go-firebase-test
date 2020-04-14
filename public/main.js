token = null

function login() {
  disableAllButtons()
  const email = document.querySelector("#email").value
  const password = document.querySelector("#password").value  
  firebase.auth().signInWithEmailAndPassword(email, password).then(res => {
    res.user.getIdToken().then(idToken => {
      token = idToken.toString()
    })
  }).then(() => enableAllButtons())
}

function ping() {
  disableAllButtons()
  fetch("/api/ping", {
    headers: {
      Authorization: token || "omit"
    }
  }).then(res => res.text()).then(text => document.querySelector("#response").innerText += text).then(() => enableAllButtons())
}

function disableAllButtons() {
  const buttons = document.querySelectorAll("button")
  for(btn of buttons) {
    btn.disabled = true
  }
}

function enableAllButtons() {
  const buttons = document.querySelectorAll("button")
  for(btn of buttons) {
    btn.disabled = false
  }
}
