/*
 * This file is part of the libvirt-go project
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 * Copyright (c) 2013 Alex Zorin
 * Copyright (C) 2016 Red Hat, Inc.
 *
 */

package libvirt

/*
#cgo pkg-config: libvirt
#include <assert.h>
#include <stdint.h>
#include "secret_events_wrapper.h"
#include "callbacks_wrapper.h"

extern void secretEventLifecycleCallback(virConnectPtr, virSecretPtr, int, int, int);
void secretEventLifecycleCallbackHelper(virConnectPtr c, virSecretPtr d,
                                     int event, int detail, void *data)
{
    secretEventLifecycleCallback(c, d, event, detail, (int)(intptr_t)data);
}

extern void secretEventGenericCallback(virConnectPtr, virSecretPtr, int);
void secretEventGenericCallbackHelper(virConnectPtr c, virSecretPtr d,
                                    void *data)
{
    secretEventGenericCallback(c, d, (int)(intptr_t)data);
}

int virConnectSecretEventRegisterAnyWrapper(virConnectPtr c,  virSecretPtr d,
                                         int eventID, virConnectSecretEventGenericCallback cb,
                                         long goCallbackId) {
    void* id = (void*)goCallbackId;
#if LIBVIR_VERSION_NUMBER < 3000000
    assert(0); // Caller should have checked version
#else
    return virConnectSecretEventRegisterAny(c, d, eventID, cb, id, freeGoCallbackHelper);
#endif
}

int virConnectSecretEventDeregisterAnyWrapper(virConnectPtr conn,
					      int callbackID)
{
#if LIBVIR_VERSION_NUMBER < 3000000
    assert(0); // Caller should have checked version
#else
    return virConnectSecretEventDeregisterAny(conn, callbackID);
#endif
}


*/
import "C"
