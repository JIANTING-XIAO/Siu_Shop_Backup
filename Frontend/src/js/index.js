import {ref} from "vue";
import axios from "axios";
export const kinds=['全部','电脑','零食','相机','旅游团']
export const selectKind=ref(0)
export const products=ref([])
export const currentPage=ref(1)
export const pageSize=ref(12)
export const totalPage=ref(1)
export const total=ref(0)
export const focusInput = () => {
    const inputElement = document.querySelector('header .SearchBox input[type="search"]');
    if (inputElement) {
        inputElement.focus();
    }
};
export const fetchProducts =async ()=>{
    const res=await axios.get("/product",{
        params:{
            page:currentPage.value,pageSize:pageSize.value
        }
    });
    products.value=res.data.product;
    total.value=res.data.total;
    totalPage.value=res.data.totalPage;
};
export const goPreviousPage = () => {
    if (currentPage.value > 1) {
        currentPage.value--;
        fetchProducts();
    }
};
export const goNextPage = () => {
    if (currentPage.value < totalPage.value) {
        currentPage.value++;
        fetchProducts();
    }
};
export const showMenu=ref(false)
