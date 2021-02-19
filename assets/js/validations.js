$(document).ready(function() { 
 
    $('#SingIn').click(function() {  
 
        $(".error").hide();
        var hasError = false;
        var emailReg = /^([\w-\.]+@([\w-]+\.)+[\w-]{2,4})?$/;
        var passreg=/^(?=.*\d)(?=.*[A-Z])(?=.*[a-z]).{8,}$/;
        var passw = $("#Password").val();
        var emailaddressVal = $("#Emailid").val();
        if(emailaddressVal == '') {
            $("#Emailid").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">*Enter email address.</span>');
            hasError = true;
        }
 
        else if(!emailReg.test(emailaddressVal)) {
            $("#Emailid").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">*Enter valid email address.</span>');
            hasError = true;
        }
 
       
        if(passw == '') {
            $("#Password").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">*Enter password.</span>');
            hasError = true;
          } 
          else if(!passreg.test(passw)) {
            $("#Password").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">*Enter valid password.</span>');
            hasError = true;
        }
          if(hasError == true) { return false; }
    });


    $('#btn_Register').click(function() {  
      
        $(".error1").hide();
        var hasError = false;
        var emailpat = /^([\w-\.]+@([\w-]+\.)+[\w-]{2,4})?$/;
        var passpat=/^(?=.*\d)(?=.*[A-Z])(?=.*[a-z]).{8,}$/;
        var mobpat=/^([6-9]{1}[0-9]{9})$/;
      
        var name = $("#Name").val();
        var mobile = $("#Mobileno").val();
        var email = $("#Emailid").val();
        var location = $("#Location").val();
        var pass = $("#Password").val();
        var cpass = $("#Cpassword").val();
        if(name == '') {
            $("#Name").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Enter first name.</span>');
            hasError = true;
          } 
          if(location == '') {
            $("#Location").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Enter location.</span>');
            hasError = true;
          } 
        if(email == '') {
            $("#Emailid").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Enter email address.</span>');
            hasError = true;
        }
 
        else if(!emailpat.test(email)) {
            $("#Emailid").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Enter valid email address.</span>');
            hasError = true;
        }
 
       
        if(pass == '') {
            $("#Password").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Set password.</span>');
            hasError = true;
          } 
          else if(!passpat.test(pass)) {
            // $("#Password").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">Set password must be in(a-z,A-Z,min 8 digit,one special($,@)).</span>');
            $("#Password").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Password should contain one capital letter and number and special characters.</span>');
            hasError = true;
        }

        if(cpass == '') {
          $("#Cpassword").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Set confirm password.</span>');
          hasError = true;
        } 
        else if(!passpat.test(cpass)) {
          $("#Cpassword").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Password should contain one capital letter and number and special characters.</span>');
          hasError = true;
      }
        else if(pass != cpass){
        $("#Cpassword").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Confirm password must be same as above</span>');
        hasError = true;
      }

        if(mobile == '') {
            $("#Mobileno").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Enter mobile no.</span>');
            hasError = true;
          } 
          else if(!mobpat.test(mobile)) {
            $("#Mobileno").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Mobile number must be 10 digit.</span>');
            hasError = true;
        }

       

          if(hasError == true) { return false; }
    });

    

});

