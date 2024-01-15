<script>
import ErrorMsg from '../components/ErrorMsg.vue';
import Msg from '../components/Msg.vue';

export default {
    components: ErrorMsg, Msg,
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

            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async setUsername() {
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
            }
            catch (e) {
                this.errormsg = "Username already used, re-try"
            }
        },
        logout() {
            localStorage.setItem("token", null)
            this.$router.push({ path: '/' });
        },
        isBanned(id) {  //implementare ban e unban
            let yes = false;
            if (this.profile.bans == null) { return false }
            for (let i = 0; i < this.profile.bans.length; i++) {

                if (this.profile.bans[i].userID == id) {
                    yes = true;
                }
            }
            return yes;
        },
        async unban(id) {
            try {
                let response = await this.$axios.get('/users/' + this.myID + "/banID/" + id,
                    {
                        headers: {
                            Authorization: this.myID
                        }
                    });
                let banID = response.data.banID
                await this.$axios.delete('/users/' + this.myID + "/bans/" + banID,
                    {
                        headers: {
                            Authorization: this.myID
                        }
                    });
                this.load()
            }
            catch (e) {
                alert("Error: " + e);
            }
        },

    },
    mounted() {
        if (!(localStorage.getItem("token") > 0)) {
            this.$router.push({ path: '/' });
            return
        }
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
                    <input type="text" class="form-control my-2" id="username" placeholder="Enter your new username"
                        v-model="this.newUsername">
                    <button type="submit" class="btn btn-primary" @click="this.setUsername()">Change</button>
                </div>
                <div class="mt-3 mb-3">
                    <h4 class=" text-danger" style="cursor:pointer" data-toggle="modal" data-target="#bansModal">
                        <svg style="width:40px;" stroke="red" :fill="'none'">
                            <use href="/feather-sprite-v4.29.0.svg#x-octagon" />
                        </svg>
                        Manage your bans
                    </h4>

                </div>
                <div class="d-flex justify-content-center">
                    <button class="btn btn-danger" @click="this.logout()">Logout</button>
                </div>
            </div>
        </div>
    </div>


    <!-- List of bans -->

    <div class="modal fade" id="bansModal" tabindex="-1" role="dialog" aria-labelledby="bansModalLabel" aria-hidden="true">
        <div class="modal-dialog" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="bansModalLabel">Bans</h5>
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                    </button>
                </div>
                <div class="modal-body">
                    <div class="container mt-5">
                        <div class="row">
                            <div class="col">


                                <div class="card mb-3" v-for="user in this.profile.bans" v-if="this.profile.bans != null" :key="user.userID">
                                    <div class="card-body d-flex align-items-center justify-content-between">
                                        <div class="me-5 p-2 bd-highlight">
                                            <h5 class="card-title">{{ user.username }}</h5>
                                        </div>
                                        <div class="p-2 bd-highlight">
                                            <button type="button" class="btn btn-outline-danger"
                                            @click="unban(user.userID)">Delete ban</button>

                                        </div>
                                    </div>
                                </div>
                                <div v-else class="d-flex align-items-center">
                                    <p>You have no bans yet</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
                </div>
            </div>
        </div>
    </div>
</template>