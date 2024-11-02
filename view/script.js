const canvas = document.getElementById("myCanvas");
const ctx = canvas.getContext("2d");

// Map to store colors for each being
const beingColors = new Map();

// Array to store previous positions for each being
const beingTrajectories = new Map();

// Function to generate a random color in hex format
function getRandomColor() {
  const letters = "0123456789ABCDEF";
  let color = "#";
  for (let i = 0; i < 6; i++) {
    color += letters[Math.floor(Math.random() * 16)];
  }
  return color;
}

// Function to draw the center point for each point
function drawCenterPoint(x, y, color) {
  ctx.beginPath();
  ctx.arc(x, y, 5, 0, Math.PI * 2);
  ctx.fillStyle = color; // Set color for the center point
  ctx.fill();
  ctx.closePath();
}

// Function to draw points with gradient fading effect
function drawPoints(data) {
  data.forEach((being, index) => {
    // Extract position and radius from the being object
    const x = being.position.X * canvas.width;
    const y = being.position.Y * canvas.height;
    const radius = being.status / 2 || 10;

    // Get or generate color for the being
    let color = beingColors.get(being.id);
    if (!color) {
      color = getRandomColor();
      beingColors.set(being.id, color);
    }

    // Debugging output
    console.log(
      `Drawing being ${
        index + 1
      }: x=${x}, y=${y}, radius=${radius}, color=${color}`
    );

    // Create a radial gradient
    const gradient = ctx.createRadialGradient(x, y, 0, x, y, radius);
    gradient.addColorStop(0, color); // Full color at the center
    gradient.addColorStop(1, "rgba(255, 255, 255, 0)"); // Transparent at the edge

    // Draw the circle with the gradient
    ctx.beginPath();
    ctx.arc(x, y, radius, 0, Math.PI * 2);
    ctx.fillStyle = gradient;
    ctx.fill();
    ctx.closePath();

    // Draw the center point for this specific point
    drawCenterPoint(x, y, color);

    // Draw trajectory
    let trajectory = beingTrajectories.get(being.id);
    if (!trajectory) {
      trajectory = [];
      beingTrajectories.set(being.id, trajectory);
    }
    trajectory.push({ x, y });
    if (trajectory.length > 3) {
      trajectory.shift(); // Keep only the last 3 positions
    }
    if (trajectory.length > 1) {
      for (let i = 1; i < trajectory.length; i++) {
        const alpha = (i + 1) / trajectory.length; // Calculate alpha value
        ctx.beginPath();
        ctx.moveTo(trajectory[i - 1].x, trajectory[i - 1].y);
        ctx.lineTo(trajectory[i].x, trajectory[i].y);
        ctx.strokeStyle = `rgba(${parseInt(color.slice(1, 3), 16)}, ${parseInt(
          color.slice(3, 5),
          16
        )}, ${parseInt(color.slice(5, 7), 16)}, ${alpha})`;
        ctx.stroke();
        ctx.closePath();
      }
    }

    // Optional: Add text labels for each point
    ctx.fillStyle = "black";
    // ctx.fillText(`Point ${index + 1}`, x + 10, y);
  });
}

// Function to fetch data from API and draw points
async function fetchDataAndDraw() {
  try {
    // Replace 'http://localhost:8080/beings' with your actual API URL
    const response = await fetch("http://localhost:8080/beings");
    const data = await response.json();

    // Debugging output
    console.log("Fetched data:", data);

    // Clear the canvas before each draw
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    // Draw points with the fetched data
    drawPoints(data);
  } catch (error) {
    console.error("Error fetching data:", error);
  }
}

// Call fetchDataAndDraw every second
setInterval(fetchDataAndDraw, 1000);
