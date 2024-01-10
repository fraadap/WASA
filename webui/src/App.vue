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
		<div class="row" v-if="$route.path != '/'"> <!-- Se non sto nella pagina iniziale carica la navbar-->
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block sidebar collapse" style="background: linear-gradient( #063ca7, #933af9);">
				<div>
					<h6 class="sidebar-heading d-flex text-center justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<h2 style="color:white;">Menu</h2>
					</h6>

					<div class="h-10 w-100 my-4">
						<RouterLink :to="'/home'" class="nav-link">
							<div class="d-flex align-items-center  ">
								<svg :fill="'none'" stroke="white" style="height:100px">
									<use href="/feather-sprite-v4.29.0.svg#home" />
								</svg>
							</div>
							<h2 style="color:white; text-align:center">Home</h2>
						</RouterLink>
					</div>

					<div class="h-25 w-100 my-4">
						<RouterLink to="/search" class="nav-link">
							<div class="d-flex align-items-center">
								<svg :fill="'none'" stroke="white" style="height:100px">
									<use href="/feather-sprite-v4.29.0.svg#search" />
								</svg>
							</div>
							<h2 style="color:white; text-align:center">Search</h2>
						</RouterLink>
					</div>

					<div class="h-25 w-100 my-4">
						<RouterLink to="/settings" class="nav-link">
							<div class="d-flex align-items-end">
								<svg :fill="'none'" stroke="white" style="height:100px">
									<use href="/feather-sprite-v4.29.0.svg#settings" />
								</svg>
							</div>
							<h2 style="color:white; text-align:center">Settings</h2>
						</RouterLink>
					</div>
					<div class="h-25 w-100 my-4">
						<RouterLink :to="'/profile/0'" class="nav-link" :us="0">
							<div class="d-flex align-items-end" >
								<svg :fill="'none'" stroke="white" style="height:100px">
									<use href="/feather-sprite-v4.29.0.svg#user" />
								</svg>
							</div>
							<h2 style="color:white; text-align:center">Profile</h2>
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
