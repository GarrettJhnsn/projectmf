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
            notify("error", "Error Submitting Request");
          }
          if(form.checkValidity()) {
            event.preventDefault();
              notify("success", "Successfully Submitted")
              setTimeout(() => form.submit(), 2000);
          }
          form.classList.add('was-validated')
        }, false)
      })
    
    }
)()

function notify(msgType, msg){
  notie.alert({
    type: msgType, // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
    text: msg,
    stay: false, // optional, default = false
    time: 3, // optional, default = 3, minimum = 1,
    position: "top" // optional, default = 'top', enum: ['top', 'bottom']
  })
}

