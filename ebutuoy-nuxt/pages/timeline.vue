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

      <!-- タイムライン -->
      <el-card class="box-card" v-for="data in timeline" :key="data.id">
        <el-row :gutter="20">
          <el-col :span="6">
            <div style="width: 120px;" class="box">
              <img :src="`https://s3-ap-northeast-1.amazonaws.com/nagisa-intern/comic/${data.comic_id}/square_thumb.jpeg`" :width="110" :height="110">
              <p>{{getUsername(data.id)}}</p>
            </div>
          </el-col>

          <el-col :span="12">
            <div style="width: 280px;" class="text item box">
              <p>{{data.content}}</p>
              <el-button type="text" class="button" @click="$router.push('/books/' + data.comic_id)">本の詳細を見る</el-button>
            </div>
          </el-col>
        </el-row>
      </el-card>
    </section>
    <!-- Footer　-->
    <div class="timeline-footer">
      <AppFooter></AppFooter>
      <nuxt/>
    </div>
  </div>
</template>

<script>
import AppHeader from '@/components/Header.vue';
import AppFooter from '@/components/Footer.vue';
import axios from 'axios'
// axios.defaults.baseURL = 'http://ebutuoy.to-hutohu.com/api/'
axios.defaults.baseURL = 'http://54.248.63.189/api/'
axios.defaults.withCredentials = true

export default {
  data() {
    return {
      radio: 1,
      activeName: 'first',
      activeIndex: '1',
      activeIndex2: '1',
      currentDate: new Date(),
      timeline: [],
      users: [],
      comics: []
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
    await axios.get('/me').catch(() => axios.post('/me', {username: 'tohutohu', password: 'tohutohu'}))
    await axios.post('/users/1/friendship').catch(() => {})
    await axios
    .get(`/timeline`)
    .then(response => (this.timeline = response.data))
    .catch(error => console.log(error))
    await axios.get('/me/follows').then(res => this.users = res.data)
    await axios.get('/comics').then(res => this.comics = res.data)
    console.log(this.timeline);
    console.log(this.users);
    console.log(this.comics);
  },
  methods: {
    getUsername(id) {
      return (this.users.find(u => u.id === id) || {username: ''}).username
    }
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

.box-card {
  width: 600px;
  height: 160px;
  margin-bottom: 10px;
  font-size: 15px;
}

.timeline-footer {
  margin-top: 20px;
}
</style>
