# PITSINKR

Open Source DNS Sinkhole

![Beta](https://shieldcn.dev/badge/status-beta-blue.svg?variant=outline&size=xs)
![License MIT](https://shieldcn.dev/badge/license-MIT-green.svg?variant=outline&size=xs)
![PRs Welcome](https://shieldcn.dev/badge/PRs-welcome-brightgreen.svg?variant=secondary&size=xs)
![Contributions Welcome](https://shieldcn.dev/badge/contributions-welcome-brightgreen.svg?variant=secondary&logo=ri%3AGoHeartFill&size=xs)

![Works on My Machine](https://shieldcn.dev/badge/works%20on-my%20machine-brightgreen.svg?size=xs)
![Tests: Eventually](https://shieldcn.dev/badge/tests-eventually-orange.svg?variant=outline&size=xs)
![Hopes & Dreams](https://shieldcn.dev/badge/Runs%20on-hopes%20%26%20dreams-FF69B4.svg?variant=secondary&size=xs)

![Go](https://shieldcn.dev/badge/Go-00ADD8.svg?logo=go&logoColor=fff&variant=branded&size=xs)
![TypeScript](https://shieldcn.dev/badge/TypeScript-3178C6.svg?logo=typescript&logoColor=fff&variant=branded&size=xs)
![React](https://shieldcn.dev/badge/React-61DAFB.svg?logo=react&logoColor=000&variant=branded&size=xs)
![Docker](https://shieldcn.dev/badge/Docker-2496ED.svg?logo=docker&logoColor=fff&variant=branded&size=xs)
![Tailwind CSS](https://shieldcn.dev/badge/Tailwind%20CSS-06B6D4.svg?logo=tailwindcss&logoColor=fff&variant=branded&size=xs)
![SQLite](https://shieldcn.dev/badge/SQLite-003B57.svg?logo=sqlite&logoColor=fff&variant=branded&size=xs)

![GitHub Contributors](https://shieldcn.dev/github/contributors/Fatuousnerd/PitSinkr.svg?variant=outline&theme=emerald&size=xs)
![GitHub Last Commit](https://shieldcn.dev/github/last-commit/Fatuousnerd/PitSinkr.svg?variant=outline&size=xs)
![GitHub Open PRs](https://shieldcn.dev/github/open-prs/Fatuousnerd/PitSinkr.svg?variant=outline&size=xs)
![GitHub Forks](https://shieldcn.dev/github/forks/Fatuousnerd/PitSinkr.svg?variant=branded&size=xs)
![GitHub Stars](https://shieldcn.dev/github/stars/Fatuousnerd/PitSinkr.svg?variant=branded&size=xs)

## DESCRIPTION

PitSinkr is a free and open source DNS sinkhole designed to block ads, trackers, malware domains, and unwanted content at the network level.

Inspired by [Pi-Hole](https://github.com/pi-hole/pi-hole), PitSinkr aims to be a modern, high-performance alternative with improved speed, easier management, and additional features. It runs on your hardware or server and filters DNS requests for your entire network.

## CURRENT STATUS

PitSinkr is in very early development stages.
Currently it can only read and log DNS traffic. Blocking functionality is not yet implemented.
We are actively working to add full sinkholing, blocklists, and more.

> You can see your traffic requests in the web dashboard available at `localhost:9056` once you start the app.

## KEY GOALS

- Drop-in replacement or better than Pi-Hole
- Excellent performance even on low-power devices
- Modern and user-friendly web interface
- Easy setup and maintenance
- Strong privacy focus
- Flexible blocklist management
- Detailed query statistics and analytics

## CURRENT CAPABILITIES

- Captures and logs DNS queries
- Basic traffic inspection
- Multi-platform support (Linux, macOS, Windows - coming soon)

## PLANNED FEATURES

- Full DNS blocking and sinkholing
- Support for popular blocklists (AdGuard, StevenBlack, etc.)
- Real-time statistics and graphs
- Whitelisting and custom rules
- Automatic blocklist updates
- DHCP server integration
- Query caching for better speed
- API for third-party integrations
- Docker and standalone installation options

## INSTALLATION (Early Stage)

> Note: Installation is currently for developers and testers only. This is not recommended for production use at the moment.

1. Clone the repository
2. Build the Go backend
3. Install dependencies for the frontend
4. Run the DNS listener (requires root/admin privileges)

Detailed setup instructions will be added as the project matures. `Docs coming soon`

## HOW TO TEST

You can currently run PitSinkr to monitor DNS traffic on your network.
Configure your router or devices to forward DNS queries to the machine running PitSinkr.

## LIMITATIONS

- No blocking yet (read-only mode only)
- Dashboard is limited
- Missing many Pi-Hole features
- Documentation is not yet available
- Not recommended for production use

We are moving fast to implement core blocking functionality.

## PRIVACY

PitSinkr keeps all your DNS data local. No telemetry or data sharing.

## CONTRIBUTING

We welcome contributions of all kinds. This is a community project.

Ways to help:

- Report bugs and issues
- Suggest features
- Submit code (Go or TypeScript)
- Improve documentation
- Test on different hardware

DEVELOPMENT
The project is split into:

- Go core for DNS server and logic
- TypeScript/React frontend

## LICENSE

This project is licensed under the MIT License.

## THANK YOU

Thank you for your interest in PitSinkr.
Stay tuned as we build a powerful next-generation DNS sinkhole.

![Made with Love](https://shieldcn.dev/badge/Made%20with-%E2%9D%A4-red.svg?size=xs)
![Sponsor This Project](https://shieldcn.dev/badge/%E2%9D%A4%EF%B8%8F%20Sponsor-this%20project-FF69B4.svg?variant=secondary&size=xs)
