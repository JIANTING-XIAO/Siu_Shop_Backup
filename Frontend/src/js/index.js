import {ref} from "vue";
import axios from "axios";
export const kinds=['全部','电脑','零食','相机','旅游团']
export const selectKind=ref(0)

export const products=ref([])
export const currentPage=ref(1)
export const pageSize=ref(12)
export const totalPage=ref(1)
export const total=ref(0)
export const searchQuery=ref('')
export const isSearching=ref(false)

export const focusInput = () => {
    const inputElement = document.querySelector('header .SearchBox input[type="search"]');
    if (inputElement) {
        inputElement.focus();
    }
};

// 防抖函数
let searchTimeout;
export const handleSearch = (query) => {
    searchQuery.value = query;
    currentPage.value = 1; // 重置到第一页
    
    // 清除之前的定时器
    if (searchTimeout) {
        clearTimeout(searchTimeout);
    }
    
    // 设置新的定时器，500ms后执行搜索
    searchTimeout = setTimeout(() => {
        fetchProducts();
    }, 500);
};

export const fetchProducts =async ()=>{
    isSearching.value = true;
    try {
        const params = {
            page: currentPage.value,
            pageSize: pageSize.value
        };
        
        // 如果有搜索关键词，添加到参数中
        if (searchQuery.value.trim()) {
            params.search = searchQuery.value.trim();
        }
        
        const res = await axios.get("/product", { params });
        products.value = res.data.product;
        total.value = res.data.total;
        totalPage.value = res.data.totalPage;
    } catch (error) {
        console.error('搜索商品时出错:', error);
    } finally {
        isSearching.value = false;
    }
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
