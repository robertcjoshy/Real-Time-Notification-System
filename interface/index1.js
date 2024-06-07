// Example notification list received from another application
const notifications = [
    { text: "Hi Guys, I'm Anna Kim, I'm from United States. I'm 25 years old Computer Engineer.", time: "1m ago" },
    { text: "Hi Guys, I'm Mark Webber, I'm from United States. I'm 25 years old Computer Engineer.", time: "5m ago" },
    // Add more notifications as needed
];

const notificationMain = document.getElementById('notification-main');
const unReadMessagesCount = document.getElementById('num-of-notif');
const markAll = document.getElementById('mark-as-read');

// Function to create a notification card
function createNotificationCard(notification) {
    const card = document.createElement('div');
    card.classList.add('notificationCard', 'unread');

    const img = document.createElement('img');
    img.alt = 'photo';
    img.src = 'githubmark.png';

    const description = document.createElement('div');
    description.classList.add('description');

    const text = document.createElement('p');
    text.innerText = notification.text;

    const time = document.createElement('p');
    time.id = 'notif-time';
    time.innerText = notification.time;

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

// Function to update unread count
function updateUnreadCount() {
    const newUnreadMessages = document.querySelectorAll('.unread');
    unReadMessagesCount.innerText = newUnreadMessages.length;
}

// Add notifications to the DOM
notifications.forEach(notification => {
    const notificationCard = createNotificationCard(notification);
    notificationMain.appendChild(notificationCard);
});

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
