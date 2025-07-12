<script setup>
import '../css/Indexstyle.css'
import {focusInput, showMenu,kinds,selectKind,products,currentPage,total,totalPage,pageSize,fetchProducts,goNextPage,goPreviousPage} from "@/js/index.js";
import {useUserStore} from "@/pinia/user.js";
import {onMounted} from "vue";
const userStore = useUserStore()
onMounted(() => {
  //获取用户信息
  if (userStore.token && !userStore.username) {
    userStore.fetchUserInfo()
  }
  //获取商品信息
  fetchProducts()
})



</script>

<template>
  <!-- 导航栏开始-->
  <header>
    <div class="left">
      <img src="/img/icon.jpg" alt="Logo" class="logo">
      <div class="logoName">秋秋商场</div>
      <nav>
        <a href="#">首页</a>
        <a href="#">商品</a>
        <a href="#">促销抢购</a>
        <a href="#">加入我们</a>
      </nav>
    </div>

    <div class="SearchBox" @click="focusInput">
      <svg class="icon" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
      <input type="search" placeholder="搜索..." />
      <kbd class="kbd kbd-sm">ctrl</kbd>
      <kbd class="kbd kbd-sm">K</kbd>
    </div>
    <!--头像开始 -->
    <div class="rightBox">
      <i class="fi fi-rr-shopping-cart"></i>
      <img :src="userStore.avatar" class="avatar1" alt="picture" @mouseenter="showMenu=true" @mouseleave="showMenu=false">
      <div class="user-menu" @mouseenter="showMenu=true" @mouseleave="showMenu=false" v-if="showMenu">
        <div class="user-info">
          <img :src="userStore.avatar" class="avatar2" alt="User Avatar">
          <span class="user-name">{{ userStore.username }}</span>
        </div>
        <div class="user-divider"></div>
        <div class="user-menu-item"><i class="fi fi-rr-user"></i>&nbsp;&nbsp;账号设置</div>
        <div class="user-menu-item"><i class="fi fi-rr-bag-shopping-minus"></i>&nbsp;&nbsp;结算订单</div>
        <div class="user-menu-item"><i class="fi fi-rr-file-minus"></i>&nbsp;&nbsp;我的订单</div>
        <div class="user-menu-item"><i class="fi fi-rr-exit"></i>&nbsp;&nbsp;退出登录</div>
      </div>
    </div>
    <!--头像结束 -->
  </header>
  <!-- 导航栏结束-->
<!--主体部分-->
  <main>
    <div class="banner">
      <div class="banner-kind">
        <div
            v-for="(item, idx) in kinds"
            :key="item"
            class="kind-item"
            :class="{ active: idx === selectKind }"
            @click="selectKind = idx"
        >
          {{ item }}
        </div>
      </div>
    </div>
    <div class="dividerLine"></div>
    <div class="fliSelect">
      <div class="fliKind">
        最近上新
      </div>
      <div class="fliKind">
        热度最高
      </div>
      <div class="fliKind">
        筛选
      </div>
    </div>
    <!--商品展示-->
    <div class="itemShow">
      <div class="item" v-for="item in products" :key="item.id">
        <div class="item-img">
          <img :src="item.product_image" alt="商品图片" />
        </div>
        <div class="item-name">
          {{ item.product_name }}
        </div>
        <div class="item-detail">
          {{ item.product_detail }}
        </div>
        <div class="item-price">
          <span>￥{{ item.price }}</span>
          <div class="item-total-sell">
            销量:{{ item.total_sell }}
          </div>
        </div>
      </div>
    </div>
    <!--商品展示-->
    <!-- 放在 main 的最下方 -->
    <div class="join">
      <button class="join-item btn" :disabled="currentPage === 1" @click="goPreviousPage">«</button>
      <span class="join-item btn btn-disabled">{{ currentPage }} / {{ totalPage }}</span>
      <button class="join-item btn" :disabled="currentPage === totalPage" @click="goNextPage">»</button>
    </div>
<!--分页结束-->

  </main>
<!--主体部分结束-->



  <footer>
    <!-- 末尾 starts  -->
    <section class="indexend">
      <div class="relative isolate flex items-center gap-x-6 overflow-hidden bg-gray-50 px-6 py-2.5 sm:px-3.5 sm:before:flex-1">
        <div class="absolute top-1/2 left-[max(-7rem,calc(50%-52rem))] -z-10 -translate-y-1/2 transform-gpu blur-2xl" aria-hidden="true">
          <div class="aspect-577/310 w-144.25 bg-linear-to-r from-[#ff80b5] to-[#9089fc] opacity-30" style="clip-path: polygon(74.8% 41.9%, 97.2% 73.2%, 100% 34.9%, 92.5% 0.4%, 87.5% 0%, 75% 28.6%, 58.5% 54.6%, 50.1% 56.8%, 46.9% 44%, 48.3% 17.4%, 24.7% 53.9%, 0% 27.9%, 11.9% 74.2%, 24.9% 54.1%, 68.6% 100%, 74.8% 41.9%)"></div>
        </div>
        <div class="absolute top-1/2 left-[max(45rem,calc(50%+8rem))] -z-10 -translate-y-1/2 transform-gpu blur-2xl" aria-hidden="true">
          <div class="aspect-577/310 w-144.25 bg-linear-to-r from-[#ff80b5] to-[#9089fc] opacity-30" style="clip-path: polygon(74.8% 41.9%, 97.2% 73.2%, 100% 34.9%, 92.5% 0.4%, 87.5% 0%, 75% 28.6%, 58.5% 54.6%, 50.1% 56.8%, 46.9% 44%, 48.3% 17.4%, 24.7% 53.9%, 0% 27.9%, 11.9% 74.2%, 24.9% 54.1%, 68.6% 100%, 74.8% 41.9%)"></div>
        </div>
        <div class="flex flex-wrap items-center gap-x-4 gap-y-2">
          <p class="text-sm/6 text-gray-900">
            <strong class="font-semibold">秋秋商场 2025</strong><svg viewBox="0 0 2 2" class="mx-2 inline size-0.5 fill-current" aria-hidden="true"><circle cx="1" cy="1" r="1" /></svg>在5-7月加入我们 会有季节优惠卡！
          </p>
          <a href="#" class="flex-none rounded-full bg-gray-900 px-3.5 py-1 text-sm font-semibold text-white shadow-xs hover:bg-gray-700 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gray-900">马上注册 <span aria-hidden="true">&rarr;</span></a>
        </div>
        <div class="flex flex-1 justify-end">
          <button type="button" class="-m-3 p-3 focus-visible:-outline-offset-4">
            <span class="sr-only">Dismiss</span>
            <svg class="size-5 text-gray-900" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" data-slot="icon">
              <path d="M6.28 5.22a.75.75 0 0 0-1.06 1.06L8.94 10l-3.72 3.72a.75.75 0 1 0 1.06 1.06L10 11.06l3.72 3.72a.75.75 0 1 0 1.06-1.06L11.06 10l3.72-3.72a.75.75 0 0 0-1.06-1.06L10 8.94 6.28 5.22Z" />
            </svg>
          </button>
        </div>
      </div>
    </section>

    <!-- 末尾 ends -->
  </footer>
</template>

<style scoped>

</style>