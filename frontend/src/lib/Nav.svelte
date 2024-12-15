<script>
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
        <a class="dropdown-item" href="/games">All games</a>
        <a class="dropdown-item" href="/platform">Per platform</a>
        <a class="dropdown-item" href="/developper">Per developper</a>
    </div>

    <input type="search" name="" id="">

    <a href="/login">Login</a>
</nav>

<style>
    .dropdown {
        position: absolute;
        display: flex;
        flex-direction: column;
        background-color: #f1f1f1;
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