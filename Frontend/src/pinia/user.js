// Frontend/pinia/user.js
import axios from "axios";
import { defineStore } from "pinia";

export const useUserStore = defineStore('user', {
    state: () => ({
        id: null,
        username: '',
        avatar: '',
        token: localStorage.getItem('token') || '', // 页面刷新自动读取
    }),
    actions: {
        setToken(token) {
            this.token = token;
            localStorage.setItem('token', token);
        },
        /*async fetchUserInfo() {
            if (!this.token) return;
            const res = await axios.get('/userinfo', {
                headers: {
                    Authorization: `Bearer ${this.token}`
                }
            });
            console.log('res.data.user:',res.data);
            this.id = res.data.user.id;
            this.username = res.data.user.name;
            this.avatar = res.data.user.avatar;
        },*/
        async fetchUserInfo() {
            if (!this.token) return;
            try {
                const res = await axios.get('/userinfo', {
                    headers: {
                        Authorization: `Bearer ${this.token}`
                    }
                });
                this.id = res.data.user.id;
                this.username = res.data.user.name;
                this.avatar = res.data.user.avatar;
            } catch (error) {
                console.error('获取用户信息失败:', error);
                this.logout(); // 如果获取失败，清除用户信息
            }
        },
        logout() {
            this.id = null;
            this.username = '';
            this.avatar = '';
            this.token = '';
            localStorage.removeItem('token');
        }
    }
});