// URL Backend Golang kita
const API_URL = "http://localhost:8080/api/portfolio";

// Fungsi utama untuk mengambil data dari Backend
async function fetchPortfolioData() {
  try {
    const response = await fetch(API_URL);

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const data = await response.json();

    // Render data ke tampilan HTML
    renderProfile(data.profile);
    renderSkills(data.skills);
    renderProjects(data.projects);
  } catch (error) {
    console.error("Gagal mengambil data dari backend:", error);
    document.getElementById("user-name").innerText = "Gagal Memuat Data";
    document.getElementById("user-about").innerText =
      "Pastikan server Golang sudah berjalan di http://localhost:8080";
  }
}

// 1. Render Section Profile / Hero
function renderProfile(profile) {
  document.getElementById("user-name").innerText = profile.name;
  document.getElementById("user-role").innerText = profile.role;
  document.getElementById("user-about").innerText = profile.about_me;
}

// 2. Render Section Skills
function renderSkills(skills) {
  const container = document.getElementById("skills-container");
  container.innerHTML = "";

  skills.forEach((skill) => {
    const skillCard = document.createElement("div");
    skillCard.className = "skill-card";
    skillCard.innerText = skill;
    container.appendChild(skillCard);
  });
}

// 3. Render Section Projects
function renderProjects(projects) {
  const container = document.getElementById("projects-container");
  container.innerHTML = "";

  projects.forEach((project) => {
    // Generate tag HTML untuk list tech_stack
    const techTagsHtml = project.tech_stack
      .map((tech) => `<span class="tech-tag">${tech}</span>`)
      .join("");

    const projectCard = document.createElement("div");
    projectCard.className = "project-card";
    projectCard.innerHTML = `
      <div>
        <h3>${project.title}</h3>
        <p>${project.description}</p>
        <div class="tech-tags">
          ${techTagsHtml}
        </div>
      </div>
      <div class="project-links">
        <a href="${project.github_url}" target="_blank">Lihat di GitHub &rarr;</a>
      </div>
    `;

    container.appendChild(projectCard);
  });
}

// Jalankan fetch saat halaman selesai di-load
document.addEventListener("DOMContentLoaded", fetchPortfolioData);
