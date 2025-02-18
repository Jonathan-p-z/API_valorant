const container=document.querySelector('.container');
const LoginLink=document.querySelectorAll('SignInLink')
const RegisterLink=document.querySelectorAll('SignUpLink')
RegisterLink.addEventListerner('click',()=>{
    container.classList.add('active');
})
LoginLink.addEventListener('click',()=>{
    container.classList.add('active');
})