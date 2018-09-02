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

      <!-- カルーセル　-->
      <!--
      <el-carousel trigger="click" height="250px" >
      <el-carousel-item v-for="item in 4" :key="item">
      <h3>{{ item }}</h3>
    </el-carousel-item>
  </el-carousel>
-->



<!-- 本一覧　-->
<el-row>
  <el-col :span="8" v-for="(o, index) in 17" :key="o" :offset="index > 0 ? 2 : 0">
    <el-card :body-style="{ padding: '0px' }">
      <!--<img :src="`/images/square_thumb${o}.png`" class="image">-->
      <img :src="`https://s3-ap-northeast-1.amazonaws.com/nagisa-intern/comic/${o}/square_thumb.jpeg`" class="image">
      <div style="padding: 14px;">
        <el-badge :value="20" class="item" :max="99">
          <span>作品タイトル</span>
        </el-badge>
        <div class="bottom clearfix">
          <span>作者名</span>
          <br>
          <div class="hashtag">
            <el-tag type="info">#ホラー</el-tag>
          </div>
          <div class="hashtag">
            <el-tag type="info">#コメディー</el-tag>
          </div>
          <div class="hashtag">
            <el-tag type="info">#人気</el-tag>
          </div>
          <el-button type="text" class="button">本の詳細を見る</el-button>
        </div>
      </div>
    </el-card>
  </el-col>
</el-row>


</section>
<!-- Footer　-->
<div class="books-footer">
  <AppFooter></AppFooter>
  <nuxt/>
</div>
</div>
</template>

<script>
import AppHeader from '@/components/Header.vue';
import AppFooter from '@/components/Footer.vue';
import axios from 'axios'

axios.defaults.withCredentials = true
// axios.defaults.baseURL = 'http://ebutuoy.to-hutohu.com/api'
axios.defaults.baseURL = 'http://54.248.63.189/api'

export default {
  data() {
    return {
      radio: 1,
      activeName: 'first',
      activeIndex: '2',
      currentDate: new Date(),
      books: []
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
  async created() {
    await axios.get('/me').catch(async () => {
      await axios.post('/me', {username: 'po', password: 'po'})
    })
    await axios
    .get(`/comics`)
    .then(response => (this.books = response.data))
    .catch(error => console.log(error))
    console.log(this.books);
  }
};
</script>

<style>
.container {

  display: flex;
  justify-content: space-around;
  align-items: center;
  flex-direction: column;
  text-align: center;
}

.el-menu-demo {
  width: 100%;
  margin-bottom: 20px;
}

.time {
  font-size: 13px;
  color: #999;
}

.item {
  margin-left: 0px;
}

.hashtag {
  display: inline-block;
  margin: 5px;
  width: auto;
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

.books-footer {
  margin-top: 20px;
}
.el-carousel__item h3 {
  color: #475669;
  font-size: 14px;
  opacity: 0.75;
  line-height: 150px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
  background-color: #d3dce6;
}
</style>
