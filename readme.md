<div align="center">
<img src="images/logo.png" width="30%">

![Yahoo! Messenger](https://img.shields.io/badge/Yahoo!%20Messenger-720E9E?style=for-the-badge&logo=yahoo&logoColor=white)
![Golang](https://img.shields.io/badge/golang-5996FF?style=for-the-badge&logo=golang&logoColor=black)

</div>

---

## [>>] About

This tool fixes the issue where **Yahoo! Web Search** does not appear in **Yahoo! Messenger** on modern versions of Windows, without requiring Compatibility Mode, by taking advantage of a Windows bug.

**(－_－) zzZ** — That's all this little tool does.

---

## [>>] Screenshots

| Before (Not Fixed) | After (Fixed) |
|:------------------:|:-------------:|
| <img src="images/Hidden.PNG" alt="Yahoo Messenger without Web Search" width="400"/> | <img src="images/fixed.PNG" alt="Yahoo Messenger with Web Search working" width="400"/> |
| *Web Search feature missing/broken* | *Web Search feature restored* |

---

## [>>] Supported Windows Versions

| Windows Version | Status |
|:----------------|:------:|
| Windows 10 22H2 | ✅ Tested & Working |
| Windows 11 Build 21996 | ✅ Tested & Working |
| Windows 11 21H2 | ✅ Tested & Working |
| Windows 11 22H2 | ✅ Tested & Working |
| Windows 11 23H2 | ✅ Tested & Working |
| Windows 11 24H2 | ✅ Tested & Working |
| Windows 11 25H2 | ⚠️ Not Tested |
| Windows 11 26H2 | ⚠️ Not Tested |

---

## [>>] How to Run

### Option 1: Pre-compiled Executable

1. **Run as Administrator** (Right-click → *Run as Administrator*)
2. **Select** `YahooMessenger.exe` when prompted
3. **Wait** for the patch to complete
4. **Press Enter** to close the window

(－_－) zzZ That's it. Launch Yahoo! Messenger and the Web Search should now be visible.

---

### Option 2: Compile from Source

```bash
# Clone the repository
git clone https://github.com/andreiixedev/YM-Search-Fixer.git
cd YM-Search-Fixer

# Initialize Go module
go mod init yahoo-fix

# Download dependencies (if any)
go mod tidy

# Build the executable
go build -o yahoo-fix.exe

# Run the compiled executable as Administrator
# (Right-click yahoo-fix.exe → Run as Administrator)
