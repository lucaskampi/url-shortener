'use client'

import React, { useState } from 'react'
import { Link, Plus, BarChart3, Copy, Check } from 'lucide-react'

export default function HomePage() {
  const [longUrl, setLongUrl] = useState('')
  const [shortCode, setShortCode] = useState('')
  const [shortUrl, setShortUrl] = useState('')
  const [loading, setLoading] = useState(false)
  const [copied, setCopied] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!longUrl) return

    setLoading(true)
    try {
      const res = await fetch('http://localhost:8080/urls', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ longUrl }),
      })
      const data = await res.json()
      setShortCode(data.shortCode)
      setShortUrl(data.shortUrl)
    } catch (err) {
      console.error('Failed to create short URL', err)
    } finally {
      setLoading(false)
    }
  }

  const handleCopy = () => {
    navigator.clipboard.writeText(shortUrl)
    setCopied(true)
    setTimeout(() => setCopied(false), 2000)
  }

  return (
    <main className="min-h-screen">
      <header className="fixed top-7 left-0 right-0 z-50">
        <div className="flex justify-center">
          <div className="pointer-events-auto w-full max-w-4xl mx-auto px-4 sm:px-6 py-4 transition-all duration-300 rounded-full bg-gray-50/50 border border-gray-300/40 backdrop-blur-md shadow-sm">
            <nav className="flex items-center justify-between">
              <a href="/" className="flex items-center space-x-2">
                <div className="bg-gradient-to-r from-blue-600 to-cyan-600 p-2 rounded-lg">
                  <Link className="w-5 h-5 sm:w-6 sm:h-6 text-white" />
                </div>
                <span className="text-lg sm:text-xl font-bold text-black">ShortURL</span>
              </a>
            </nav>
          </div>
        </div>
      </header>

      <div className="max-w-4xl mx-auto p-6 pt-40">
        <div className="text-center mb-12">
          <h1 className="text-4xl sm:text-5xl font-bold text-gray-900 mb-4">
            Shorten Your Links
          </h1>
          <p className="text-lg text-gray-600">
            Fast, simple, and free. Create short links in seconds.
          </p>
        </div>

        <div className="bg-white rounded-2xl shadow-lg p-6 sm:p-8 border border-gray-200">
          <form onSubmit={handleSubmit} className="space-y-4">
            <div className="flex flex-col sm:flex-row gap-3">
              <input
                type="url"
                placeholder="Paste your long URL here..."
                value={longUrl}
                onChange={(e) => setLongUrl(e.target.value)}
                className="flex-1 px-4 py-3 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                required
              />
              <button
                type="submit"
                disabled={loading}
                className="px-6 py-3 bg-black hover:bg-gray-900 text-white font-medium rounded-xl flex items-center justify-center gap-2 transition-colors disabled:opacity-50"
              >
                {loading ? (
                  'Creating...'
                ) : (
                  <>
                    <Plus className="w-5 h-5" />
                    Shorten
                  </>
                )}
              </button>
            </div>
          </form>

          {shortUrl && (
            <div className="mt-6 p-4 bg-gray-50 rounded-xl border border-gray-200">
              <p className="text-sm text-gray-500 mb-2">Your shortened URL:</p>
              <div className="flex items-center gap-3">
                <a
                  href={shortUrl}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-blue-600 hover:text-blue-800 font-medium text-lg"
                >
                  {shortUrl}
                </a>
                <button
                  onClick={handleCopy}
                  className="p-2 hover:bg-gray-200 rounded-lg transition-colors"
                  title="Copy to clipboard"
                >
                  {copied ? (
                    <Check className="w-5 h-5 text-green-600" />
                  ) : (
                    <Copy className="w-5 h-5 text-gray-600" />
                  )}
                </button>
              </div>
              <div className="mt-3 flex gap-4">
                <a
                  href={`/stats/${shortCode}`}
                  className="flex items-center gap-2 text-sm text-gray-600 hover:text-gray-900"
                >
                  <BarChart3 className="w-4 h-4" />
                  View Stats
                </a>
              </div>
            </div>
          )}
        </div>
      </div>
    </main>
  )
}