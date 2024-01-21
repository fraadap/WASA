<script>
import { RouterLink } from 'vue-router';
import LoadingSpinner from '../components/LoadingSpinner.vue'

export default {
    components: LoadingSpinner,
    data: function () {
        return {
            text: "",
            users: [],
            cacheUsers: [],
            loading: false,
            inputSearch: "",
        }
    },
    methods: {
        async search() {
            if (this.inputSearch == "") { this.users = []; return }
            if (this.cacheUsers[this.inputSearch] == null) {
                this.loading = true
                try {
                    let response = await this.$axios.get('/search/' + this.inputSearch,
                        {
                            headers: {
                                Authorization: parseInt(localStorage.getItem("token"))
                            }
                        });
                    this.users = response.data
                    if (response.data == null) {
                        this.cacheUsers[this.inputSearch] = []
                    } else {
                        this.cacheUsers[this.inputSearch] = response.data
                    }
                }
                catch (e) {
                    alert(e);
                }
                this.loading = false
            } else {
                this.users = this.cacheUsers[this.inputSearch]
            }


        },
    },
    mounted() {
        if (!(localStorage.getItem("token") > 0)) {
            this.$router.push({ path: '/' });
            return
        }
    }
}

</script>   

<template>
    <div class="container mx-4 my-5">

        <input type="text" id="searchInput" v-model="inputSearch" v-on:keyup="this.search()" class="form-control mb-4"
            placeholder="Cerca utenti..." style="width:20%">
        <div v-if="users != null">
            <div class="card user-card my-3" v-for="u in users" :key="u.user.userID">
                <RouterLink :to="'/profile/' + u.user.userID" class="nav-link" :us="u.user.userID">
                    <div class="card-body">
                        <div class="container mt-3 text-center">
                            <div class="row">
                                <div class="col fs-3">
                                    <p>{{ u.user.username }}</p>
                                </div>
                                <div class="col fs-5">
                                    <p>Photos</p>
                                    <p>{{ u.nPhotos }}</p>
                                </div>
                                <div class="col fs-5">
                                    <p>Followers</p>
                                    <p>{{ u.followers ? u.followers.length : 0 }}</p>
                                </div>
                                <div class="col fs-5">
                                    <p>Followings</p>
                                    <p>{{ u.followings ? u.followings.length : 0 }}</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </RouterLink>
            </div>
        </div>
        <LoadingSpinner :loading="this.loading"></LoadingSpinner>
    </div>
</template>