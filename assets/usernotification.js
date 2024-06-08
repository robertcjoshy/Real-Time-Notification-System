
function fetchNotifications() {
    fetch('/user/notification')
        .then(response => response.json())
        .then(data => {
            if (data.message && Array.isArray(data.message)) {
                // Display the notifications
                console.log(data.message);
                // Update the DOM to show the notifications
                data.message.forEach(notification => {
                    const notificationCard = createNotificationCard(notification);
                    notificationMain.insertBefore(notificationCard, notificationMain.firstChild);
                });
                // Update the unread messages count
                updateUnreadCount();
            }
            // Immediately call fetchNotifications again to poll
            fetchNotifications();
        })
        .catch(error => {
            console.error('Error fetching notifications:', error);
            // Retry immediately after an error
            fetchNotifications();
        });
}

// Function to create a notification card
function createNotificationCard(notification) {
    const card = document.createElement('div');
    card.classList.add('notificationCard', 'unread');

    const img = document.createElement('img');
    img.alt = 'photo';
    img.src = '/assets/githubmark.png'; // Correct path to the static image

    const description = document.createElement('div');
    description.classList.add('description');

    const text = document.createElement('p');
    text.innerText = notification;

    const time = document.createElement('p');
    time.classList.add('notif-time');
    time.innerText = getCurrentTime();

    description.appendChild(text);
    description.appendChild(time);
    card.appendChild(img);
    card.appendChild(description);

    card.addEventListener('click', () => {
        card.classList.remove('unread');
        updateUnreadCount();
    });

    return card;
}

// Function to get the current time in 24-hour format (HH:MM)
function getCurrentTime() {
    const now = new Date();
    const hours = String(now.getHours()).padStart(2, '0');
    const minutes = String(now.getMinutes()).padStart(2, '0');
    return `${hours}:${minutes}`;
}

// Function to update unread count
function updateUnreadCount() {
    const newUnreadMessages = document.querySelectorAll('.unread');
    unReadMessagesCount.innerText = newUnreadMessages.length;
}

const notificationMain = document.getElementById('notification-main');
const unReadMessagesCount = document.getElementById('num-of-notif');
const markAll = document.getElementById('mark-as-read');
const logoutButton = document.getElementById('logout-button');


// Initialize the unread messages count
updateUnreadCount();

// Mark all as read functionality
markAll.addEventListener('click', () => {
    const allMessages = document.querySelectorAll('.notificationCard');
    allMessages.forEach(message => {
        message.classList.remove('unread');
    });
    updateUnreadCount();
});
// logout functionality
logoutButton.addEventListener('click', () => {
    fetch('/user/logout', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (response.redirected) {
            window.location.href = response.url;
        } else {
            console.error('Logout failed');
        }
    })
    .catch(error => {
        console.error('Error during logout:', error);
    });
});

document.addEventListener('DOMContentLoaded', function() {
    // Start polling for notifications
    fetchNotifications();
});


