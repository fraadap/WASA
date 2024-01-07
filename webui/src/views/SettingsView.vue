<script>
import ErrorMsg from '../components/ErrorMsg.vue';
import Msg from '../components/Msg.vue';
import { RouterLink } from 'vue-router';

export default {
    components: ErrorMsg,Msg,
    data: function () {
        return {
            errormsg: null,
            msg: null,
            profile: {},
            myID: parseInt(localStorage.getItem("token")),
            newUsername: "",

        }
    },
    methods: {
        async load() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.get('/users/' + this.myID,
                    { headers: { Authorization: this.myID } });
                this.profile = response.data;
                console.log(this.profile)

            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async setUsername() {
            console.log(this.newUsername)
            try {
                let response = await this.$axios.put('/users/' + this.myID,
                    {
                        userID: this.myID,
                        username: this.newUsername,
                    },
                    {
                        headers: { Authorization: this.myID },
                    },
                );
                this.msg = "Username changed succesfully"
                console.log(response)
            }
            catch (e) {
                this.errormsg = "Username already used, re-try"
            }
        },
        logout() {
            console.log("logout")
        },
        async unban() {

        }

    },
    mounted() {
        this.load()
    }
}
</script>

<template>
    <ErrorMsg v-if="this.errormsg" :msg="this.errormsg"></ErrorMsg>
    <Msg v-if="this.msg" :msg="this.msg"></Msg>
    <div class="container h-100">
        <div class="row h-100 justify-content-center align-items-center p-3">
            <div class="col-6">
                <h2 class="text-center">Settings</h2>
                <div class="mt-3 card my-5 p-4">
                    <h3 class="text-center mb-3">Change your username</h3>
                    <label for="username">New username:</label>
                    <input type="text" class="form-control" id="username" placeholder="Enter your new username"
                        v-model="this.newUsername">
                    <button type="submit" class="btn btn-primary" @click="this.setUsername()">Change</button>
                </div>
                <div class="mt-3 mb-3">
                    <h4 class=" text-danger" style="cursor:pointer">
                        <svg style="width:40px;" stroke="red" :fill="'none'">
                            <use href="/feather-sprite-v4.29.0.svg#x-octagon" />
                        </svg>
                        Manage your bans
                    </h4>

                </div>
                <div class="d-flex justify-content-center">
                    <button class="btn btn-danger">Logout</button>
                </div>
            </div>
        </div>
    </div>
</template>