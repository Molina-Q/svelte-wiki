<script>
	import { Link } from "svelte-routing";
    let isOpen = $state(false);

    function toggleDropdown() {
        isOpen = !isOpen;
    }
</script>

<nav>
    <!-- <img src="" alt=""> -->
    <a href="/">Home</a>

    <button type="button" onclick={toggleDropdown}>Games</button>

    <div class={`dropdown ${!isOpen && 'dropdown-closed'}`}>
        <Link class="dropdown-item" to="/games">All games</Link>
        <Link class="dropdown-item" to="/platform">Per platform</Link>
        <Link class="dropdown-item" to="/developper">Per developper</Link>
        <Link class="dropdown-item" to="/create/games">Add game</Link>
    </div>

    <input type="search" name="" id="">

    <a href="/login">Login</a>
</nav>

<style>
    nav {
        padding: 1rem;
        background-color: var(--svg-white);

        display: flex;
        justify-content: space-between;
        border-bottom: 3px solid var(--svg-blue);
    }
    
    .dropdown {
        position: absolute;
        display: flex;
        flex-direction: column;
        background-color: var(--svg-white);

        color: #333;
        text-decoration: none;
    }

    .dropdown-closed {
        display: none;
    }

    .dropdown-item {
        position: absolute;
        display: block;
        padding: 1rem;
        background-color: #f1f1f1;
        color: #333;
        text-decoration: none;

        &:hover {
            background-color: #333;
            color: #f1f1f1;
        }
    }
</style>