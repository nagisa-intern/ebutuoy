<template>
  <div>
    <!-- Header　-->
    <div>
      <AppHeader></AppHeader>
      <nuxt/>
    </div>
    <!-- タブ　-->
    <section class="container">
      <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
        <el-menu-item index="1"><a href="./timeline" target="_self">タイムライン</a></el-menu-item>
        <el-menu-item index="2"><a href="./books" target="_self">本から選ぶ</a></el-menu-item>
        <el-submenu index="3">
          <template slot="title">アカウント</template>
          <el-menu-item index="2-1"><a href="./users/1" target="_self">プロフィール</a></el-menu-item>
          <el-menu-item index="2-2">設定</el-menu-item>
          <el-menu-item index="2-3">ログアウト</el-menu-item>
        </el-submenu>
      </el-menu>

      <el-row>
        <el-col :span="8" v-for="(o, index) in 2" :key="o" :offset="index > 0 ? 2 : 0">
          <el-card :body-style="{ padding: '0px' }">
            <img src="~/assets/images/hamburger.png" class="image">
            <div style="padding: 14px;">
              <span>Yummy hamburger</span>
              <div class="bottom clearfix">
                <time class="time">{{ currentDate }}</time>
                <el-button type="text" class="button">本の詳細を見る</el-button>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>

    </section>
    <!-- Footer　-->
    <div>
      <AppFooter></AppFooter>
      <nuxt/>
    </div>
  </div>
</template>

<script>
import AppHeader from '@/components/Header.vue';
import AppFooter from '@/components/Footer.vue';

export default {
  data() {
    return {
      radio: 1,
      activeName: 'first',
      activeIndex: '2',
      currentDate: new Date()
    };
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    }
  },
  components: {
    AppHeader,
    AppFooter,
  },
};
</script>

<style>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: space-around;
  align-items: center;
  flex-direction: column;
  text-align: center;
}
.time {
  font-size: 13px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
}

.button {
  padding: 0;
  float: right;
}

.image {
  width: 100%;
  display: block;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both
}
</style>
