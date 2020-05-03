// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from "vue";
import App from "./App";
import router from "./router";
import firebase from "firebase";

Vue.config.productionTip = false;

const config = {
  apiKey: "AIzaSyBg3fAAKVE2wVAqet0N8Bvac0dc6uRgPNQ",
  authDomain: "barcelona-zoo.firebaseapp.com",
  databaseURL: "https://barcelona-zoo.firebaseio.com",
  projectId: "barcelona-zoo",
  storageBucket: "barcelona-zoo.appspot.com",
  messagingSenderId: "161724186936",
  appId: "1:161724186936:web:81fe189fb255f8daa6c1cd"
};
firebase.initializeApp(config);

/* eslint-disable no-new */
new Vue({
  el: "#app",
  router,
  components: { App },
  template: "<App/>"
});
