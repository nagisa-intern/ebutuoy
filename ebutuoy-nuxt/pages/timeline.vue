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
      <el-card class="box-card">
        <!--
        <div slot="header" class="clearfix">
        <span>◯◯◯◯◯(本タイトル)</span>
        <el-button style="float: right; padding: 3px 0" type="text">本の詳細画面へ</el-button>
      </div>
    -->
    <el-row :gutter="20">

      <el-col :span="6">
        <div style="width: 120px;" class="box">
          <img src="~/assets/images/sample_icon.jpeg" class="image"  :width="120" :height="120">
          <p>user_name</p>
        </div>
      </el-col>

      <el-col :span="12">
        <div style="width: 280px;" class="text item box">
          <p>ああああああああああ</p>
          <el-button type="text" class="button">本の詳細を見る</el-button>
        </div>
      </el-col>
    </el-row>
  </el-card>


  <el-card class="box-card">
    <!--
    <div slot="header" class="clearfix">
    <span>◯◯◯◯◯(本タイトル)</span>
    <el-button style="float: right; padding: 3px 0" type="text">本の詳細画面へ</el-button>
  </div>
-->
<el-row :gutter="20">

  <el-col :span="6">
    <div style="width: 120px;" class="box">
      <img src="~/assets/images/sample_icon.jpeg" class="image"  :width="120" :height="120">
      <p>user_name</p>
    </div>
  </el-col>

  <el-col :span="12">
    <div style="width: 280px;" class="text item box">
      <p>ああああああああああ</p>
      <el-button type="text" class="button">本の詳細を見る</el-button>
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

export default {
  data() {
    return {
      radio: 1,
      activeName: 'first',
      activeIndex: '1',
      activeIndex2: '1',
      currentDate: new Date(),
      timeline: []
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
    await axios
    .get(`http://ebutuoy.to-hutohu.com/api/timeline`)
    .then(response => (this.timeline = response))
    .catch(error => console.log(error))
    console.log(this.timeline);
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
