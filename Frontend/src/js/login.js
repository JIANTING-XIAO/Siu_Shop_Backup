import { ref } from 'vue'
import axios from "axios";
import {useUserStore} from "@/pinia/user.js";
const userStore = useUserStore()

const container = ref(null)
const passwordInput = ref('')
const accountInput = ref('')
const telephoneInput = ref('')
const errorMessage = ref('')

const toggleRegister = () => {
    if (container.value) {
        container.value.classList.add('active')
    }
    // 清空输入框
    accountInput.value = ''
    passwordInput.value = ''
    errorMessage.value = ''
}

const toggleLogin = () => {
    if (container.value) {
        container.value.classList.remove('active')
    }
    // 清空输入框
    accountInput.value = ''
    passwordInput.value = ''
    telephoneInput.value = ''
    errorMessage.value = ''
}

async function loginPost(){
    errorMessage.value = ''

    if (!accountInput.value || !passwordInput.value) {
        errorMessage.value = '用户名或密码不能为空'
        return
    }

    try {
        const response = await axios.post('/login', {
            account: accountInput.value,
            password: passwordInput.value
        }, {
            headers: {
                'Content-Type': 'application/json'
            }
        })

        if(response.data.code===200){
            //登录成功
            console.log('登录成功:',response.data.token.token)
            //保存token到localStorage
            if (response.data.token){
                localStorage.setItem('token', response.data.token.token);
            }
            userStore.setToken(response.data.token.token)
            await userStore.fetchUserInfo() // 获取用户信息
            console.log('用户信息:', userStore.id, userStore.username, userStore.avatar)
            //添加登录成功的跳转逻辑
            window.location.href = '/'
        }else {
            //登录失败
            errorMessage.value = response.data.message || '登录失败'
            //可以在这里显示错误信息给用户
        }
    } catch (error) {
        console.error('登录时发生错误：', error)
        errorMessage.value = error.response?.data?.message|| '登录失败，请稍后重试'

    }
}
async function registerPost(){
    try{
        const response=await axios.post('/register',{
            account:accountInput.value,
            password:passwordInput.value,
            telephone:telephoneInput.value
        },{
            headers: {
                'Content-Type': 'application/json'
            }
        })
        if (response.data.code === 200) {
            // 注册成功
            console.log('注册成功：', response.data)
            // 保存token
            if (response.data.token) {
                localStorage.setItem('token', response.data.token.token)
            }
        } else {
            errorMessage.value=response.data.message|| '注册失败'
        }
    }catch (error) {
        console.error('注册时发生错误：', error)
        errorMessage.value = error.response?.data?.message|| '注册失败，请稍后重试'
    }
}
export { container, toggleRegister, toggleLogin ,loginPost,registerPost,errorMessage,telephoneInput, accountInput, passwordInput }