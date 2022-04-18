export default function getEnv(name) {
    return window?.configs?.[name] || process.env[name]
  }