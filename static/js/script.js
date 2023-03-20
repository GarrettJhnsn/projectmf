// Example Bootstrap starter JavaScript for disabling form submissions if there are invalid fields
(function () {
    'use strict'
  
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    var forms = document.querySelectorAll('.needs-validation')
  
    // Loop over them and prevent submission
    Array.prototype.slice.call(forms)
      .forEach(function (form) {
        form.addEventListener('submit', function (event) {
          if (!form.checkValidity()) {
            event.preventDefault()
            event.stopPropagation()
          }
          form.classList.add('was-validated')
        }, false)
      })
    }
)()

document.getElementById("formSubmit").addEventListener("click", function() {
    notify("Error Submitting Request", "error")
})

// Notie Alert Function
function notify(msg, msgType) {
    notie.alert({
        type: msgType,
        text: msg,
        stay: false,
        time: 3,
        position: "top"
    })
}