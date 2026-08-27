import { useState, useEffect } from 'react'

const BASE_URL = 'http://localhost:8080'

function App() {
  // Asosiy state-lar
  const [users, setUsers] = useState([])
  const [formData, setFormData] = useState({ id: '', name: '', age: '', email: '' })
  const [editingId, setEditingId] = useState(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)
  const [logs, setLogs] = useState([])

  // HTTP Request Log-ga yozish yordamchi funksiyasi
  const addLog = (method, path, status) => {
    const newLog = {
      id: Date.now() + Math.random(),
      method,
      path,
      status,
      time: new Date().toLocaleTimeString()
    }
    setLogs(prev => [newLog, ...prev])
  }

  // 1. GET /users - barcha userlarni yuklash
  const fetchUsers = async () => {
    setLoading(true)
    setError(null)
    try {
      const res = await fetch(`${BASE_URL}/users`)
      addLog('GET', '/users', res.status)

      if (!res.ok) {
        throw new Error(`Server xatosi: ${res.status} ${res.statusText}`)
      }

      const data = await res.json()
      setUsers(Array.isArray(data) ? data : [])
    } catch (err) {
      console.error('Fetch users error:', err)
      setError(err.message || "Foydalanuvchilarni yuklashda xatolik yuz berdi")
      if (!err.message?.includes('Server xatosi:')) {
        addLog('GET', '/users', 'ERR')
      }
    } finally {
      setLoading(false)
    }
  }

  // Sahifa birinchi marta ochilganda userlarni yuklash
  useEffect(() => {
    fetchUsers()
  }, [])

  // Forma inputlari o'zgarganda state-ni yangilash
  const handleInputChange = (e) => {
    const { name, value } = e.target
    setFormData(prev => ({ ...prev, [name]: value }))
  }

  // 2. POST /users (Create) yoki PATCH /users/{id} (Update)
  const handleSubmit = async (e) => {
    e.preventDefault()

    if (!formData.name.trim() || !formData.email.trim() || !formData.age || (!editingId && !formData.id)) {
      alert("Iltimos, barcha maydonlarni to'ldiring!")
      return
    }

    setLoading(true)
    setError(null)

    if (editingId) {
      // PATCH /users/{id}
      const path = `/users/${editingId}`
      const payload = {
        name: formData.name.trim(),
        age: parseInt(formData.age, 10),
        email: formData.email.trim()
      }

      try {
        const res = await fetch(`${BASE_URL}${path}`, {
          method: 'PATCH',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(payload)
        })
        addLog('PATCH', path, res.status)

        if (!res.ok) {
          throw new Error(`Yangilashda xatolik: ${res.status} ${res.statusText}`)
        }

        // Formani tozalash va tahrirlash holatidan chiqish
        setFormData({ id: '', name: '', age: '', email: '' })
        setEditingId(null)
        await fetchUsers()
      } catch (err) {
        console.error('Update error:', err)
        setError(err.message)
        if (!err.message?.includes('Yangilashda xatolik:')) {
          addLog('PATCH', path, 'ERR')
        }
      } finally {
        setLoading(false)
      }
    } else {
      // POST /users (ID bilan birga)
      const path = '/users'
      const payload = {
        id: parseInt(formData.id, 10),
        name: formData.name.trim(),
        age: parseInt(formData.age, 10),
        email: formData.email.trim()
      }

      try {
        const res = await fetch(`${BASE_URL}${path}`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(payload)
        })
        addLog('POST', path, res.status)

        if (!res.ok) {
          throw new Error(`Yaratishda xatolik: ${res.status} ${res.statusText}`)
        }

        // Formani tozalash
        setFormData({ id: '', name: '', age: '', email: '' })
        await fetchUsers()
      } catch (err) {
        console.error('Create error:', err)
        setError(err.message)
        if (!err.message?.includes('Yaratishda xatolik:')) {
          addLog('POST', path, 'ERR')
        }
      } finally {
        setLoading(false)
      }
    }
  }

  // 3. Edit tugmasi bosilganda formani to'ldirish
  const handleEdit = (user) => {
    setFormData({
      id: user.id !== undefined ? user.id : '',
      name: user.name || '',
      age: user.age !== undefined ? user.age : '',
      email: user.email || ''
    })
    setEditingId(user.id)
    setError(null)
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }

  // Edit rejimini bekor qilish
  const handleCancelEdit = () => {
    setFormData({ id: '', name: '', age: '', email: '' })
    setEditingId(null)
  }

  // 4. DELETE /users/{id}
  const handleDelete = async (id) => {
    const isConfirmed = window.confirm(`Haqiqatan ham ID: ${id} bo'lgan foydalanuvchini o'chirmoqchimisiz?`)
    if (!isConfirmed) return

    setLoading(true)
    setError(null)
    const path = `/users/${id}`

    try {
      const res = await fetch(`${BASE_URL}${path}`, {
        method: 'DELETE'
      })
      addLog('DELETE', path, res.status)

      if (!res.ok) {
        throw new Error(`O'chirishda xatolik: ${res.status} ${res.statusText}`)
      }

      // Agar tahrirlanayotgan user o'chirilsa, tahrirlashni bekor qilish
      if (editingId === id) {
        handleCancelEdit()
      }

      await fetchUsers()
    } catch (err) {
      console.error('Delete error:', err)
      setError(err.message)
      if (!err.message?.includes("O'chirishda xatolik:")) {
        addLog('DELETE', path, 'ERR')
      }
    } finally {
      setLoading(false)
    }
  }

  // Status kodi uchun rang badge klassini aniqlash
  const getBadgeClass = (status) => {
    if (typeof status === 'number') {
      if (status >= 200 && status < 300) return 'badge-2xx'
      if (status >= 300 && status < 400) return 'badge-3xx'
      if (status >= 400 && status < 500) return 'badge-4xx'
      if (status >= 500) return 'badge-5xx'
    }
    return 'badge-err'
  }

  return (
    <div className="container">
      {/* Header */}
      <header className="header">
        <h1>REST API Demo</h1>
        <p className="subtitle">User Profile CRUD &bull; Golang Backend (<code>net/http</code>)</p>
      </header>

      {/* Xatolik xabari */}
      {error && (
        <div className="alert-error">
          <span>⚠️ {error}</span>
          <button onClick={() => setError(null)} className="btn-close">&times;</button>
        </div>
      )}

      {/* Asosiy kontent: Chapda Form, O'ngda Userlar ro'yxati */}
      <div className="main-grid">
        {/* Chap panel: Forma */}
        <section className="card-section form-section">
          <h2>{editingId ? 'Edit User' : 'Create User'}</h2>
          <form onSubmit={handleSubmit} className="user-form">
            <div className="form-group">
              <label htmlFor="id">ID</label>
              <input
                id="id"
                type="number"
                name="id"
                value={formData.id}
                onChange={handleInputChange}
                placeholder="Masalan: 1"
                disabled={!!editingId}
                required={!editingId}
              />
            </div>

            <div className="form-group">
              <label htmlFor="name">Name</label>
              <input
                id="name"
                type="text"
                name="name"
                value={formData.name}
                onChange={handleInputChange}
                placeholder="Masalan: Ali"
                required
              />
            </div>

            <div className="form-group">
              <label htmlFor="age">Age</label>
              <input
                id="age"
                type="number"
                name="age"
                value={formData.age}
                onChange={handleInputChange}
                placeholder="Masalan: 22"
                min="1"
                max="120"
                required
              />
            </div>

            <div className="form-group">
              <label htmlFor="email">Email</label>
              <input
                id="email"
                type="email"
                name="email"
                value={formData.email}
                onChange={handleInputChange}
                placeholder="Masalan: ali@mail.com"
                required
              />
            </div>

            <div className="form-buttons">
              {!editingId ? (
                <button type="submit" className="btn btn-create" disabled={loading}>
                  {loading ? 'Kutilmoqda...' : 'Create User'}
                </button>
              ) : (
                <>
                  <button type="submit" className="btn btn-update" disabled={loading}>
                    {loading ? 'Kutilmoqda...' : 'Update User'}
                  </button>
                  <button
                    type="button"
                    onClick={handleCancelEdit}
                    className="btn btn-cancel"
                    disabled={loading}
                  >
                    Cancel Edit
                  </button>
                </>
              )}
            </div>
          </form>
        </section>

        {/* O'ng panel: Userlar ro'yxati */}
        <section className="card-section list-section">
          <div className="list-header">
            <h2>Userlar ro'yxati ({users.length})</h2>
            <button onClick={fetchUsers} className="btn-refresh" disabled={loading}>
              🔄 Refresh
            </button>
          </div>

          {loading && users.length === 0 ? (
            <div className="loading-state">Yuklanmoqda...</div>
          ) : users.length === 0 ? (
            <div className="empty-state">
              Hozircha foydalanuvchilar yo'q. Chapdagi forma orqali yangi user qo'shing.
            </div>
          ) : (
            <div className="user-cards">
              {users.map(user => (
                <div key={user.id} className={`user-card ${editingId === user.id ? 'is-editing' : ''}`}>
                  <div className="user-card-body">
                    <div className="user-id-badge">ID: {user.id}</div>
                    <div className="user-field">
                      <span className="field-label">Name:</span>
                      <span className="field-value highlight">{user.name}</span>
                    </div>
                    <div className="user-field">
                      <span className="field-label">Age:</span>
                      <span className="field-value">{user.age} yosh</span>
                    </div>
                    <div className="user-field">
                      <span className="field-label">Email:</span>
                      <span className="field-value">{user.email}</span>
                    </div>
                  </div>

                  <div className="user-card-actions">
                    <button
                      onClick={() => handleEdit(user)}
                      className="btn btn-sm btn-edit"
                      disabled={loading}
                    >
                      Edit
                    </button>
                    <button
                      onClick={() => handleDelete(user.id)}
                      className="btn btn-sm btn-delete"
                      disabled={loading}
                    >
                      Delete
                    </button>
                  </div>
                </div>
              ))}
            </div>
          )}
        </section>
      </div>

      {/* Pastki panel: HTTP Request Log */}
      <section className="card-section log-panel">
        <div className="log-header">
          <h2>HTTP Request Log</h2>
          {logs.length > 0 && (
            <button onClick={() => setLogs([])} className="btn-clear-logs">
              Clear Logs
            </button>
          )}
        </div>

        {logs.length === 0 ? (
          <div className="log-empty">Hali hech qanday HTTP so'rov amalga oshirilmadi.</div>
        ) : (
          <div className="log-list">
            {logs.map(log => (
              <div key={log.id} className="log-item">
                <span className="log-time">{log.time}</span>
                <span className={`log-method method-${log.method.toLowerCase()}`}>
                  {log.method}
                </span>
                <span className="log-path">{log.path}</span>
                <span className="log-arrow">&rarr;</span>
                <span className={`badge ${getBadgeClass(log.status)}`}>
                  {log.status}
                </span>
              </div>
            ))}
          </div>
        )}
      </section>
    </div>
  )
}

export default App
