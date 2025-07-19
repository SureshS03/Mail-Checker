# Email Checker

A lightweight and efficient email verification tool written in Go. Perfect for integrating into backend authentication systems to validate user emails during registration and login processes.

## Overview

This tool performs comprehensive email validation through multiple verification layers, making it ideal for backend integration where you need to ensure email authenticity before user registration or account recovery.

## Features

- **Syntax Validation**: RFC-compliant email format checking
- **MX Record Verification**: Validates domain mail exchange records
- **Data Breach Detection**: Checks against public breach databases via RapidAPI
- **SMTP Verification**: Real-time mailbox existence validation
- **Backend Ready**: Designed for integration into authentication systems
- **Lightweight**: Single binary with minimal dependencies

## Compatibility & Limitations

### **Fully Supported**
- **Gmail** (gmail.com) - Complete verification support
- **Custom Business Domains** - Works with most corporate email systems
- **Most Email Providers** - Generally reliable verification

### **Limited Support**
- **Yahoo Mail** (yahoo.com, ymail.com) - May have SMTP verification issues
- **Outlook/Hotmail** (outlook.com, hotmail.com, live.com) - Restricted SMTP access

> **Important**: Yahoo and Outlook implement strict SMTP policies that often block external verification attempts. Consider implementing fallback logic for these providers.

## Infrastructure Requirements

### Server Prerequisites
Your server must meet these requirements for reliable SMTP verification:

- **Clean Static IP**: Use a static IP address with good reputation
- **IP Not Blacklisted**: Ensure your IP isn't listed on spam databases
- **Reverse DNS Setup**: Configure PTR records for better deliverability
- **Port 25 Access**: Outbound SMTP access required

### Check Your IP Reputation
Before deployment, verify your server's reputation:
- [Spamhaus Lookup](https://www.spamhaus.org/lookup/)
- [MXToolbox Blacklist Check](https://mxtoolbox.com/blacklists.aspx)
- [Postmaster Tools](https://postmaster.google.com/)

> **Note**: If your IP is blacklisted, SMTP verification will fail. Consider using a VPS with clean IP reputation or email verification services.

## Installation & Setup

### Prerequisites
- Go 1.20 or higher
- RapidAPI account for breach detection
- Server with clean IP reputation

### Quick Start
```bash
# Clone or download the code
git clone <repository-url>
cd mail-checker

# Set your RapidAPI key
export API="your-rapidapi-key-here"

# Build the application
go build -o email-verifier .

# Run verification
./email-verifier user@example.com
```

## API Configuration

### RapidAPI Setup
1. Sign up at [RapidAPI](https://rapidapi.com/)
2. Subscribe to [Breach Directory API](https://rapidapi.com/rohan-patra/api/breachdirectory)
3. Get your API key from the dashboard

### Rate Limits (Important)
- **Free Plan**: Only **10 requests per month**
- **Basic Plan**: 1,000+ requests (paid)

> **Production Warning**: The free plan's 10 requests/month limit makes it unsuitable for production use. Consider upgrading or implementing caching strategies.

## Usage Examples

### Basic Verification
```bash
# Single email check
./email-verifier john.doe@gmail.com

# Expected output:
# Email syntax looks good
# Domain is exits
# Checking any data breach for: john.doe@gmail.com
# No data breach found for this email.
# validator@gmail.com
# Email exists: john.doe@gmail.com
```


## Error Codes & Handling

The tool returns different exit codes and messages:

### Success Cases
- `Email syntax looks good` - Syntax validation passed
- `Domain is exits` - MX records found
- `Email exists: <email>` - SMTP verification successful

### Error Cases
- `SMTP dial failed` - Cannot connect to mail server
- `HELO failed` - SMTP handshake failed
- `MAIL FROM failed` - Sender validation failed
- `Mail is not Exits` - Email address doesn't exist

## Production Considerations

### For Backend Integration
1. **Implement Caching**: Cache verification results to avoid repeated API calls
2. **Handle Rate Limits**: Monitor RapidAPI usage and implement fallbacks
3. **Async Processing**: Run verification in background for better UX
4. **Timeout Handling**: Set appropriate timeouts for SMTP connections
5. **Fallback Logic**: Have backup validation for Yahoo/Outlook domains


## Troubleshooting

### Common Issues

**SMTP Verification Fails**
```bash
Error: SMTP dial failed: dial tcp: connect: connection refused
```
- Check if your server's IP is blacklisted
- Verify port 25 access
- Try from a different IP/server


**Yahoo/Outlook Always Fails**
- This is expected behavior due to their restrictions
- Implement alternative validation for these domains
- Consider fallback to syntax + MX record validation only

## Development Notes

This tool is specifically designed for backend integration into authentication systems. Key integration points:

- **User Registration**: Validate emails before creating accounts
- **Password Recovery**: Verify email exists before sending reset links  
- **Email Updates**: Validate new email addresses in profile updates
- **Bulk Validation**: Process email lists for marketing campaigns

## Security Considerations

- **Rate Limiting**: Implement rate limiting to prevent abuse
- **Input Validation**: Always sanitize email inputs
- **Logging**: Log verification attempts for security monitoring
- **Privacy**: Be mindful of data breach information exposure

## Contributing

Feel free to submit issues and enhancement requests. When contributing:
1. Test with multiple email providers
2. Consider rate limiting implications
3. Update documentation for any new features

## License

MIT License - feel free to use in your projects.

---

**Perfect for backend email validation in registration systems, contact forms, and user authentication flows.**