<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {}
</script>

<template>
	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASA Photo</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse"
			data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>

	<div class="container-fluid">
		<div class="row" v-if="this.$route.path != '/'"> <!-- Se non sto nella pagina iniziale carica la navbar-->
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<h6
						class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>General</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink :to="'/home'" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#home" />
								</svg>
								Home
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/search" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#search" />
								</svg>
								Search
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/settings" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#settings" />
								</svg>
								Settings
							</RouterLink>
						</li>
					</ul>

					<h6
						class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>Secondary menu</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink to="/home" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#file-text" />
								</svg>
								Item 1
							</RouterLink>
						</li>
					</ul>
					<div class="position-absolute bottom-0 start-50 translate-middle-x w-100">
						<RouterLink :to="'/profile/0'" class="nav-link" :us="0">
							<div class="d-flex align-items-end">
								<div style="color:white" class="bg-dark fs-1"> <svg style="height:25%; width:25%;"
										fill="#ffffff">
										<use href="/feather-sprite-v4.29.0.svg#user" />
									</svg>
									<span class="ms-3 mb-0">Profile</span>
								</div>
							</div>
						</RouterLink>
					</div>
				</div>
			</nav>

			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<RouterView />
			</main>
		</div>
		<div v-else> <!-- ALtrimenti carica solo la parte di login -->
			<RouterView />
		</div>
	</div>
</template>

<style></style>
