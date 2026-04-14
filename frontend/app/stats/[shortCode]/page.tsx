'use client'

import React, { useEffect, useState } from 'react'
import { useParams } from 'next/navigation'
import { Link, ArrowLeft, BarChart3, MousePointer, Calendar } from 'lucide-react'

interface Stats {
  shortCode: string
  clickCount: number
  createdAt: string
}

export default function StatsPage() {
  const params = useParams()
  const shortCode = params.shortCode as string
  const [stats, setStats] = useState<Stats | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const res = await fetch(`http://localhost:8080/urls/${shortCode}/stats`)
        if (!res.ok) throw new Error('Failed to fetch stats')
        const data = await res.json()
        setStats(data)
      } catch (err) {
        setError('Could not load stats')
      } finally {
        setLoading(false)
      }
    }

    fetchStats()
  }, [shortCode])

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
              <a
                href="/"
                className="flex items-center gap-2 text-gray-600 hover:text-gray-900"
              >
                <ArrowLeft className="w-4 h-4" />
                Back
              </a>
            </nav>
          </div>
        </div>
      </header>

      <div className="max-w-4xl mx-auto p-6 pt-40">
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900">URL Statistics</h1>
          <p className="text-gray-600 mt-2">View performance metrics for your shortened link</p>
        </div>

        {loading && (
          <div className="bg-white rounded-2xl shadow-lg p-8 border border-gray-200 text-center">
            <div className="animate-pulse">Loading stats...</div>
          </div>
        )}

        {error && (
          <div className="bg-white rounded-2xl shadow-lg p-8 border border-gray-200 text-center">
            <p className="text-red-600">{error}</p>
            <a href="/" className="text-blue-600 hover:underline mt-4 inline-block">
              Create a new short URL
            </a>
          </div>
        )}

        {stats && !loading && !error && (
          <div className="bg-white rounded-2xl shadow-lg p-6 sm:p-8 border border-gray-200">
            <div className="grid grid-cols-1 sm:grid-cols-3 gap-6">
              <div className="p-6 bg-gray-50 rounded-xl border border-gray-200">
                <div className="flex items-center gap-3 mb-4">
                  <div className="p-2 bg-blue-100 rounded-lg">
                    <MousePointer className="w-5 h-5 text-blue-600" />
                  </div>
                  <span className="text-gray-600 font-medium">Total Clicks</span>
                </div>
                <p className="text-4xl font-bold text-gray-900">{stats.clickCount}</p>
              </div>

              <div className="p-6 bg-gray-50 rounded-xl border border-gray-200">
                <div className="flex items-center gap-3 mb-4">
                  <div className="p-2 bg-blue-100 rounded-lg">
                    <BarChart3 className="w-5 h-5 text-blue-600" />
                  </div>
                  <span className="text-gray-600 font-medium">Short Code</span>
                </div>
                <p className="text-2xl font-mono text-gray-900">{stats.shortCode}</p>
              </div>

              <div className="p-6 bg-gray-50 rounded-xl border border-gray-200">
                <div className="flex items-center gap-3 mb-4">
                  <div className="p-2 bg-blue-100 rounded-lg">
                    <Calendar className="w-5 h-5 text-blue-600" />
                  </div>
                  <span className="text-gray-600 font-medium">Created</span>
                </div>
                <p className="text-lg text-gray-900">
                  {new Date(stats.createdAt).toLocaleDateString()}
                </p>
              </div>
            </div>

            <div className="mt-6 pt-6 border-t border-gray-200">
              <p className="text-sm text-gray-500 mb-2">Original URL:</p>
              <a
                href={`http://localhost:8080/${shortCode}`}
                target="_blank"
                rel="noopener noreferrer"
                className="text-blue-600 hover:text-blue-800 break-all"
              >
                http://localhost:8080/{shortCode}
              </a>
            </div>
          </div>
        )}
      </div>
    </main>
  )
}