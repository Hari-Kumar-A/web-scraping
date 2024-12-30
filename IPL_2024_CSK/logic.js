async function fetchMatches() {
  try {
    const response = await fetch("http://localhost:8080/matches");
    if (!response.ok) {
      throw new Error("Failed to fetch data");
    }
    const matches = await response.json();

    // Filter matches for "CHENNAI SUPER KINGS"
    const filteredMatches = matches.filter((match) =>
      match.title.toUpperCase().includes("CHENNAI SUPER KINGS")
    );

    displayMatches(filteredMatches);

    // Calculate and display the win percentage
    calculateWinPercentage(filteredMatches);
  } catch (error) {
    document.getElementById("errorMessage").innerText =
      "Error: " + error.message;
  }
}

function displayMatches(matches) {
  const container = document.getElementById("matchesContainer");
  container.innerHTML = ""; // Clear any existing content

  if (matches.length === 0) {
    container.innerHTML =
      '<p>No matches involving "CHENNAI SUPER KINGS" found.</p>';
    return;
  }

  matches.forEach((match, index) => {
    const matchElement = document.createElement("div");
    matchElement.classList.add("match");

    // Determine if CSK won or lost
    if (match.result.toUpperCase().includes("CHENNAI SUPER KINGS WON")) {
      matchElement.classList.add("win"); // Add green background for win
    } else {
      matchElement.classList.add("lose"); // Add red background for loss
    }

    matchElement.innerHTML = `
                    <h2>Match ${index + 1}: ${match.title}</h2>
                    <p><strong>Stadium:</strong> ${match.stadium}</p>
                    <p><strong>Result:</strong> ${match.result}</p>
                `;
    container.appendChild(matchElement);
  });
}

function calculateWinPercentage(matches) {
  const totalMatches = matches.length;
  const wins = matches.filter((match) =>
    match.result.toUpperCase().includes("CHENNAI SUPER KINGS WON")
  ).length;

  const winPercentage = ((wins / 14) * 100).toFixed(2); // Calculate percentage based on 14 matches
  const statsContainer = document.getElementById("stats");
  statsContainer.innerText = `Chennai Super Kings Win Percentage: ${winPercentage}% (${wins} wins out of 14 matches)`;
}

// Fetch matches on page load
window.onload = fetchMatches;
