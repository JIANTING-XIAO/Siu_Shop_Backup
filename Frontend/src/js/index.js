import {ref} from "vue";
export const kinds=['全部','电脑','零食','相机','旅游团']
export const selectKind=ref(0)
export const products=ref([])
export const focusInput = () => {
    const inputElement = document.querySelector('header .SearchBox input[type="search"]');
    if (inputElement) {
        inputElement.focus();
    }
};

export const showMenu=ref(false)
